package service

import (
	"context"

	"github.com/sichangSun/course-enrollment-manager/domain/model"
	"github.com/sichangSun/course-enrollment-manager/domain/repository"
)

// CourseService ...
type CourseService struct {
	CourseRepository repository.CourseRepository
}

// NewCourseService ...
func NewCourseService(courseRepo repository.CourseRepository) *CourseService {
	return &CourseService{CourseRepository: courseRepo}
}

// GetAllCoursesOutput ...
type GetAllCoursesOutput struct {
	Courses []*model.CourseWithMulitSchedules
}

// GetAllCourses ...
func (cs *CourseService) GetAllCourses(ctx context.Context) (*GetAllCoursesOutput, error) {
	rows, err := cs.CourseRepository.GetAllCourses(ctx)
	if err != nil {
		return nil, err
	}

	courseMap := make(map[int]*model.CourseWithMulitSchedules)
	for _, detail := range rows {
		if course, exists := courseMap[detail.ID]; exists {
			course.Schedules = append(course.Schedules, &model.CourseSchedule{
				DayOfWeek: detail.DayOfWeek,
				Period:    detail.Period,
			})
		} else {
			courseMap[detail.ID] = &model.CourseWithMulitSchedules{
				ID:           detail.ID,
				CourseName:   detail.CourseName,
				Semester:     detail.Semester,
				InstructorID: detail.Instructor,
				Credits:      detail.Credits,
				TeacherName:  detail.TeacherName,
				Position:     detail.Position,
				Schedules: []*model.CourseSchedule{
					{
						DayOfWeek: detail.DayOfWeek,
						Period:    detail.Period,
					},
				},
			}
		}
	}
	var courses []*model.CourseWithMulitSchedules
	for _, course := range courseMap {
		courses = append(courses, course)
	}
	return &GetAllCoursesOutput{Courses: courses}, nil
}
