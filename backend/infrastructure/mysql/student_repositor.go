package mysql

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/sichangSun/course-enrollment-manager/domain/model"
	"github.com/sichangSun/course-enrollment-manager/domain/repository"
)

// studentRepository ...
type StudentRepository struct {
	db *sqlx.DB
}

// NewCourseRepository ...
func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{db}
}

// GetStudentByUserEmail
func (s *StudentRepository) GetStudentByUserEmail(con context.Context, email string) (*model.Student, error) {
	query := `SELECT id, name, password,email
	          FROM students WHERE email = ?;`

	var student model.Student
	err := s.db.Get(&student, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrStudentNotFound
		}
		return nil, err
	}
	return &student, nil
}

// GetStudentByID
func (s *StudentRepository) GetStudentByID(ctx context.Context, ID int) (*model.Student, error) {
	query := `SELECT id, name, password,email
	          FROM students WHERE id = ?;`

	var student model.Student
	err := s.db.Get(&student, query, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrStudentNotFound
		}
		return nil, err
	}
	return &student, nil

}

// UpdatePassword
func (s *StudentRepository) UpdatePassword(ctx context.Context, stu *model.Student) error {
	query := `UPDATE students
						SET password = ?
						WHERE id = ?;`

	_, err := s.db.Exec(query, stu.Password, stu.ID)
	if err != nil {
		return err
	}
	return nil

}

// GetStudentCourses
func (s *StudentRepository) GetStudentCourses(ctx context.Context, studentID int) ([]*model.StudentCoursesDetail, error) {
	query := `
	SELECT
		stu.id AS student_id,
		c.id AS course_id,
		c.name AS course_name,
		c.semester,
		c.instructor,
		c.credits,
		t.name AS teacher_name,
		t.position,
		cs.day_of_week,
		cs.period
	FROM students stu
	JOIN students_courses sc ON stu.id = sc.student_id
	JOIN courses c ON sc.course_id = c.id
	JOIN teachers t ON c.instructor = t.id
	JOIN course_schedules cs ON c.id = cs.course_id
	WHERE stu.id = ?;`

	var courses []*model.StudentCoursesDetail
	if err := s.db.Select(&courses, query, studentID); err != nil {
		log.Fatalf("Error querying the database: %v", err)
	}

	return courses, nil
}

// RegisterCourse
func (s *StudentRepository) RegisterCourse(ctx context.Context, sc *model.StudentCourse) error {
	query := `INSERT INTO students_courses
		(student_id, course_id) VALUES
		(:student_id, :course_id);`

	_, err := s.db.NamedExecContext(ctx, query, sc)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
	}
	return nil
}
