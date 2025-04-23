package tests

import (
	"api-repository/internal/services/user-service/service/internal/handlers"
	userservice "api-repository/pkg/api/user-service"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"testing"
)

func setupTest(t *testing.T) (*handlers.AuthHandler, sqlmock.Sqlmock, context.Context) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db: %v", err)
	}

	secretKey := "test-secret-key"
	authHandler := handlers.NewAuthHandler(db, secretKey)
	ctx := context.Background()

	return authHandler, mock, ctx
}

func TestRegister_Success(t *testing.T) {
	authHandler, mock, ctx := setupTest(t)

	// Исправлен SQL-запрос (VALUES -> values)
	mock.ExpectExec(regexp.QuoteMeta(`INSERT into users (id, username, email, password_hash) values ($1, $2, $3, $4)`)).
		WithArgs(
			sqlmock.AnyArg(),
			"testuser",
			"test@example.com",
			sqlmock.AnyArg(),
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &userservice.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		PasswordConfirm: "password123",
	}

	resp, err := authHandler.Register(ctx, req)

	// Добавлена проверка ошибки перед проверкой resp
	if !assert.NoError(t, err) {
		return
	}

	assert.NotEmpty(t, resp.Token)
	assert.NotEmpty(t, resp.Uuid)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRegister_PasswordMismatch(t *testing.T) {
	authHandler, _, ctx := setupTest(t)

	req := &userservice.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		PasswordConfirm: "password456", // Несовпадающий пароль
	}

	resp, err := authHandler.Register(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)

	statusErr, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, statusErr.Code())
	assert.Equal(t, "passwords do not match", statusErr.Message())
}

func TestRegister_DuplicateEmail(t *testing.T) {
	authHandler, mock, ctx := setupTest(t)

	// Создаем ошибку дублирования
	pqErr := &pq.Error{
		Code:    "23505",
		Message: "duplicate key value violates unique constraint",
	}

	// Исправленный SQL-запрос
	mock.ExpectExec(regexp.QuoteMeta(`INSERT into users (id, username, email, password_hash) values ($1, $2, $3, $4)`)).
		WithArgs(
			sqlmock.AnyArg(),
			"testuser",
			"test@example.com",
			sqlmock.AnyArg(),
		).
		WillReturnError(pqErr)

	req := &userservice.RegisterRequest{
		Username:        "testuser",
		Email:           "test@example.com",
		Password:        "password123",
		PasswordConfirm: "password123",
	}

	resp, err := authHandler.Register(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user with this email already exists")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLogin_Success(t *testing.T) {
	authHandler, mock, ctx := setupTest(t)

	// Создаем хеш пароля для сравнения
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	// Настраиваем ожидаемый запрос к базе данных
	rows := sqlmock.NewRows([]string{"id", "password_hash"}).
		AddRow("user-uuid-123", string(passwordHash))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password_hash FROM users WHERE email = $1`)).
		WithArgs("test@example.com").
		WillReturnRows(rows)

	req := &userservice.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	resp, err := authHandler.Login(ctx, req)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Token)
	assert.Equal(t, "user-uuid-123", resp.Uuid)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLogin_UserNotFound(t *testing.T) {
	authHandler, mock, ctx := setupTest(t)

	// Настраиваем ожидаемый запрос к базе данных, который вернет ошибку "no rows"
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password_hash FROM users WHERE email = $1`)).
		WithArgs("nonexistent@example.com").
		WillReturnError(sql.ErrNoRows)

	req := &userservice.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	resp, err := authHandler.Login(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)

	statusErr, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.NotFound, statusErr.Code())
	assert.Equal(t, "user not found", statusErr.Message())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLogin_IncorrectPassword(t *testing.T) {
	authHandler, mock, ctx := setupTest(t)

	// Создаем хеш пароля для сравнения
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	// Настраиваем ожидаемый запрос к базе данных
	rows := sqlmock.NewRows([]string{"id", "password_hash"}).
		AddRow("user-uuid-123", string(passwordHash))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password_hash FROM users WHERE email = $1`)).
		WithArgs("test@example.com").
		WillReturnRows(rows)

	req := &userservice.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword", // Неправильный пароль
	}

	resp, err := authHandler.Login(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, resp)

	statusErr, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unauthenticated, statusErr.Code())
	assert.Equal(t, "incorrect password", statusErr.Message())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
