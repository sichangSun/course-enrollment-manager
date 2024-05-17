package repository

import (
	"context"

	"github.com/sichangSun/course-enrollment-manager/domain/model"
)

// CourseRepository is ...
type CourseRepository interface {
	// GetAllCourses ...
	GetAllCourses(ctx context.Context) ([]*model.CourseScheduleDetail, error)
	//GetGetOneCourseDetail
	GetGetOneCourseDetail(ctx context.Context, courseID int) ([]*model.CourseScheduleDetail, error)
}
