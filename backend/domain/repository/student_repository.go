package repository

import (
	"context"
	"errors"

	"github.com/sichangSun/course-enrollment-manager/domain/model"
)

var ErrStudentNotFound = errors.New("user not found")

type StudentRepository interface {
	// GetStudentByUserEmail for loginin
	GetStudentByUserEmail(ctx context.Context, email string) (*model.Student, error)
	// GetStudentByID
	GetStudentByID(ctx context.Context, ID int) (*model.Student, error)
	// UpdatePassword
	UpdatePassword(ctx context.Context, stu *model.Student) error
	// GetStudentCourses
	GetStudentCourses(ctx context.Context, studentID int) ([]*model.StudentCoursesDetail, error)
	// RegisterCourse
	RegisterCourse(ctx context.Context, sc *model.StudentCourse) error
	// UnRegisterCourse
	UnRegisterCourse(ctx context.Context, sc *model.StudentCourse) error
}
