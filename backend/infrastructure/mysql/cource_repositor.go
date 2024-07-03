package mysql

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/sichangSun/course-enrollment-manager/domain/model"
)

// CourseRepository ...
type CourseRepository struct {
	db *sqlx.DB
}

// NewCourseRepository ...
func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db}
}

// GetAllCourses  ...
func (r *CourseRepository) GetAllCourses(ctx context.Context) ([]*model.CourseScheduleDetail, error) {
	query := `
	SELECT
		c.id,
		c.name AS course_name,
		c.semester,
		c.instructor,
		c.credits,
		t.name AS teacher_name,
		t.position,
		cs.day_of_week,
		cs.period
	FROM courses c
	JOIN teachers t ON c.instructor = t.id
	JOIN course_schedules cs ON c.id = cs.course_id
	ORDER BY c.id ASC;`

	var courses []*model.CourseScheduleDetail
	if err := r.db.Select(&courses, query); err != nil {
		log.Fatalf("Error querying the database: %v", err)
	}

	return courses, nil
}

// GetGetOneCourseDetail ..
func (r *CourseRepository) GetGetOneCourseDetail(ctx context.Context, courseID int) ([]*model.CourseScheduleDetail, error) {
	query := `
	SELECT
		c.id,
		c.name AS course_name,
		c.semester,
		c.instructor,
		c.credits,
		t.name AS teacher_name,
		t.position,
		cs.day_of_week,
		cs.period,
		c.description
	FROM courses c
	JOIN teachers t ON c.instructor = t.id
	JOIN course_schedules cs ON c.id = cs.course_id
	WHERE c.id = ?`

	var courses []*model.CourseScheduleDetail
	if err := r.db.Select(&courses, query, courseID); err != nil {
		log.Fatalf("Error querying the database: %v", err)
	}

	return courses, nil
}
