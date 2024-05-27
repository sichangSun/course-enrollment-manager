package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sichangSun/course-enrollment-manager/domain/model"
	"github.com/sichangSun/course-enrollment-manager/domain/repository"
)

var ErrInvalidPassword = errors.New("invalid password")

type StudentService struct {
	StudentRepository repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) *StudentService {
	return &StudentService{StudentRepository: studentRepo}
}

// LoginInput ...
type LoginInput struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// LoginOutput ...
type LoginOutput struct {
	Student *model.Student
}

// UpdatePasswordInput ...
type UpdatePasswordInput struct {
	StudentID   int    `validate:"required"`
	OldPassword string `validate:"required"`
	NewPassword string `validate:"required,min=8,max=20"`
}

// GetStudentCoursesOutput
type GetStudentCoursesOutput struct {
	StudentCourses []*model.StudentCoursesWithMulitSchedules
}

// RegisterCourseInput
type RegisterCourseInput struct {
	StudentID int
	CourseID  int
}

type UnRegisterCourseInput struct {
	StudentID int
	CourseID  int
}

// Loginin
func (s *StudentService) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	if err := validator.New().Struct(input); err != nil {
		return nil, err
	}
	student, err := s.StudentRepository.GetStudentByUserEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if err := model.CompareHashAndPassword(student.Password, input.Password); err != nil {
		return nil, ErrInvalidPassword
	}

	return &LoginOutput{Student: student}, nil
}

// UpdatePassword
func (s *StudentService) UpdatePassword(ctx context.Context, input *UpdatePasswordInput) error {
	if err := validator.New().Struct(input); err != nil {
		return err
	}
	student, err := s.StudentRepository.GetStudentByID(ctx, input.StudentID)
	if err != nil {
		return err
	}

	if err := model.CompareHashAndPassword(student.Password, input.OldPassword); err != nil {
		return ErrInvalidPassword
	}
	password, err := model.PasswordEncrypt(input.NewPassword)
	if err != nil {
		return err
	}

	newS := &model.Student{
		ID:       input.StudentID,
		Password: password,
	}

	if err := s.StudentRepository.UpdatePassword(ctx, newS); err != nil {
		return err
	}

	return nil
}

// GetStudentCourses
func (s *StudentService) GetStudentCourses(ctx context.Context, studentID int) (*GetStudentCoursesOutput, error) {
	rows, err := s.StudentRepository.GetStudentCourses(ctx, studentID)
	if len(rows) == 0 {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	courseMap := make(map[int]*model.StudentCoursesWithMulitSchedules)
	for _, detail := range rows {
		if course, exists := courseMap[detail.CourseID]; exists {
			course.Schedules = append(course.Schedules, &model.CourseSchedule{
				DayOfWeek: detail.DayOfWeek,
				Period:    detail.Period,
			})
		} else {
			courseMap[detail.CourseID] = &model.StudentCoursesWithMulitSchedules{
				StudentID:    detail.StudentID,
				CourseID:     detail.CourseID,
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
	var courses []*model.StudentCoursesWithMulitSchedules
	for _, course := range courseMap {
		courses = append(courses, course)
	}
	return &GetStudentCoursesOutput{StudentCourses: courses}, nil

}

// RegisterCourse
func (s *StudentService) RegisterCourse(ctx context.Context, input *RegisterCourseInput) error {
	studentCourse := &model.StudentCourse{
		StudentID: input.StudentID,
		CourseID:  input.CourseID,
	}
	err := s.StudentRepository.RegisterCourse(ctx, studentCourse)
	if err != nil {
		return err
	}
	return nil
}

// UnRegisterCourse
func (s *StudentService) UnRegisterCourse(ctx context.Context, input *UnRegisterCourseInput) error {
	studentCourse := &model.StudentCourse{
		StudentID: input.StudentID,
		CourseID:  input.CourseID,
	}
	err := s.StudentRepository.UnRegisterCourse(ctx, studentCourse)
	if err != nil {
		return err
	}
	return nil
}
