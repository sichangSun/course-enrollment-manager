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
