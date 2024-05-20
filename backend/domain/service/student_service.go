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
