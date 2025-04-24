package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func InsertLesson(pg *sql.DB, req *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	id := uuid.New()

	_, err := pg.Exec("INSERT INTO lessons (id, section_id, name, description) VALUES ($1, $2, $3, $4)",
		id,
		req.SectionId,
		req.Name,
		req.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("pgInsertLesson: failed to insert lesson into database: %v", err)
	}

	return &textService.CreateLessonResponse{
		Id: id.String(),
	}, nil
}

func SelectLesson(pg *sql.DB, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	query := `
		SELECT
			l.id AS lesson_id,
			l.section_id AS section_id,
			l.name AS lesson_name,
			l.description AS lesson_description,
			array_agg(DISTINCT v.id) FILTER (WHERE v.id IS NOT NULL) AS video_ids,
			array_agg(DISTINCT f.id) FILTER (WHERE f.id IS NOT NULL) AS file_ids,
			array_agg(DISTINCT e.id) FILTER (WHERE e.id IS NOT NULL) AS exercise_ids,
			array_agg(DISTINCT c.id) FILTER (WHERE c.id IS NOT NULL) AS comment_ids,
			l.rating AS lesson_rating
		FROM
			lessons l
		LEFT JOIN
			videos v ON l.id = v.lesson_id
		LEFT JOIN
			files f ON l.id = f.lesson_id
		LEFT JOIN
			exercises e ON l.id = e.lesson_id
		LEFT JOIN
			comments c ON l.id = c.lesson_id
		WHERE
			l.id = $1
		GROUP BY
			l.id, l.section_id, l.name, l.description, l.rating;
	`

	lessonRow := pg.QueryRow(query, req.Id)

	var (
		id, sectionId, name, description           string
		videoIds, fileIds, exerciseIds, commentIds pq.StringArray
		rating                                     int32
	)

	err := lessonRow.Scan(&id, &sectionId, &name, &description,
		&videoIds, &fileIds, &exerciseIds, &commentIds, &rating)
	if err != nil {
		return nil, fmt.Errorf("pgSelectLesson: failed to scan lesson: %v", err)
	}

	lesson := &textService.Lesson{
		Id:          id,
		SectionId:   sectionId,
		Name:        name,
		Description: description,
		VideoIds:    videoIds,
		FileIds:     fileIds,
		ExerciseIds: exerciseIds,
		CommentIds:  commentIds,
		Rating:      rating,
	}

	return &textService.GetLessonResponse{
		Lesson: lesson,
	}, nil
}

func SelectLessons(pg *sql.DB) (*textService.GetLessonsResponse, error) {
	lessons := make([]*textService.Lesson, 0, 10)

	query := `
		SELECT
			l.id AS lesson_id,
			l.section_id AS section_id,
			l.name AS lesson_name,
			l.description AS lesson_description,
			array_agg(DISTINCT v.id) FILTER (WHERE v.id IS NOT NULL) AS video_ids,
			array_agg(DISTINCT f.id) FILTER (WHERE f.id IS NOT NULL) AS file_ids,
			array_agg(DISTINCT e.id) FILTER (WHERE e.id IS NOT NULL) AS exercise_ids,
			array_agg(DISTINCT c.id) FILTER (WHERE c.id IS NOT NULL) AS comment_ids,
			l.rating AS lesson_rating
		FROM
			lessons l
		LEFT JOIN
			videos v ON l.id = v.lesson_id
		LEFT JOIN
			files f ON l.id = f.lesson_id
		LEFT JOIN
			exercises e ON l.id = e.lesson_id
		LEFT JOIN
			comments c ON l.id = c.lesson_id
		GROUP BY
			l.id, l.section_id, l.name, l.description, l.rating;
	`

	lessonRows, err := pg.Query(query)
	if err != nil {
		return nil, fmt.Errorf("pgSelectLessons: failed to query lessons: %v", err)
	}
	defer lessonRows.Close()

	for lessonRows.Next() {
		var (
			id, sectionId, name, description           string
			videoIds, fileIds, exerciseIds, commentIds pq.StringArray
			rating                                     int32
		)

		err := lessonRows.Scan(&id, &sectionId, &name, &description,
			&videoIds, &fileIds, &exerciseIds, &commentIds, &rating)
		if err != nil {
			return nil, fmt.Errorf("pgSelectLessons: failed to scan lesson: %v", err)
		}

		lesson := &textService.Lesson{
			Id:          id,
			SectionId:   sectionId,
			Name:        name,
			Description: description,
			VideoIds:    videoIds,
			FileIds:     fileIds,
			ExerciseIds: exerciseIds,
			CommentIds:  commentIds,
			Rating:      rating,
		}

		lessons = append(lessons, lesson)
	}

	if err := lessonRows.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectLessons: failed to iterate over lesson rows: %v", err)
	}

	return &textService.GetLessonsResponse{
		Lessons: lessons,
	}, nil
}

func IncreaseRating(pg *sql.DB, req *textService.IncreaseRatingRequest) (*textService.IncreaseRatingResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET rating = rating + 1 WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgIncreaseRating: failed to increase lesson rating: %v", err)
	}

	return &textService.IncreaseRatingResponse{
		Id: req.Id,
	}, nil
}

func DecreaseRating(pg *sql.DB, req *textService.DecreaseRatingRequest) (*textService.DecreaseRatingResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET rating = rating - 1 WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDecreaseRating: failed to decrease lesson rating: %v", err)
	}

	return &textService.DecreaseRatingResponse{
		Id: req.Id,
	}, nil
}

func DeleteLesson(pg *sql.DB, req *textService.DeleteLessonRequest) (*textService.DeleteLessonResponse, error) {
	_, err := pg.Exec("DELETE FROM lessons WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDeleteLesson: failed to delete lesson from database: %v", err)
	}

	return &textService.DeleteLessonResponse{
		Id: req.Id,
	}, nil
}
