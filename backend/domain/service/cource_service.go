package service

import (
	"context"
	"errors"

	"github.com/sichangSun/course-enrollment-manager/domain/model"
	"github.com/sichangSun/course-enrollment-manager/domain/repository"
)

var ErrNotFound = errors.New("course not found")

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

// GetOneCoursesDetailOutput
type GetOneCoursesDetailOutput struct {
	Courses *model.CourseWithMulitSchedules
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
				CourseID:     detail.ID,
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

// GetOneCourseDetail(ctx,courseID)
func (cs *CourseService) GetGetOneCourseDetail(ctx context.Context, courseID int) (*GetOneCoursesDetailOutput, error) {
	rows, err := cs.CourseRepository.GetGetOneCourseDetail(ctx, courseID)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, ErrNotFound
	}
	var schedules []*model.CourseSchedule
	for _, raw := range rows {
		schedule := &model.CourseSchedule{
			DayOfWeek: raw.DayOfWeek,
			Period:    raw.Period,
		}
		schedules = append(schedules, schedule)
	}
	courseDetail := &model.CourseWithMulitSchedules{
		CourseID:     rows[0].ID,
		CourseName:   rows[0].CourseName,
		Semester:     rows[0].Semester,
		InstructorID: rows[0].Instructor,
		Credits:      rows[0].Credits,
		TeacherName:  rows[0].TeacherName,
		Position:     rows[0].Position,
		Description:  rows[0].Description,
		Schedules:    schedules,
	}
	return &GetOneCoursesDetailOutput{Courses: courseDetail}, nil
}
