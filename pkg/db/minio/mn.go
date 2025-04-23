package minio

import (
	"api-repository/pkg/utils"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type MnConfig struct {
	Endpoint       string `yaml:"MINIO_ENDPOINT" env:"MINIO_ENDPOINT" env-default:"localhost:9000"`
	AccessKey      string `yaml:"MINIO_ACCESS_KEY" env:"MINIO_ACCESS_KEY" env-default:"minioadmin"`
	SecretKey      string `yaml:"MINIO_SECRET_KEY" env:"MINIO_SECRET_KEY" env-default:"minioadmin"`
	Region         string `yaml:"MINIO_REGION" env:"MINIO_REGION" env-default:"us-east-1"`
	UseSSL         bool   `yaml:"MINIO_USE_SSL" env:"MINIO_USE_SSL" env-default:"false"`
	ForcePathStyle bool   `yaml:"MINIO_FORCE_PATH_STYLE" env:"MINIO_FORCE_PATH_STYLE" env-default:"true"`
}

func NewVideoMinioConnection(cfg MnConfig) *minio.Client {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		utils.GetSugaredLogger().Fatalf("can't connect to minio | err: %v", err)
		return nil
	}

	createBucket(client, "videos")

	return client
}

func createBucket(client *minio.Client, bucketName string) {
	ctx := context.Background()

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("ошибка проверки бакета: %v", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
			Region: "us-east-1",
		})
		if err != nil {
			log.Fatalf("ошибка создания бакета: %v", err)
		}
		log.Printf("Бакет %s создан", bucketName)
	} else {
		log.Printf("Бакет %s уже существует", bucketName)
	}
}
