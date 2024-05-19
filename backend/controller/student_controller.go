package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sichangSun/course-enrollment-manager/domain/repository"
	"github.com/sichangSun/course-enrollment-manager/domain/service"
	"github.com/sichangSun/course-enrollment-manager/session"
)

// StudentController..
type StudentController struct {
	StudentService *service.StudentService
}

// NewStudentController..
func NewStudentController(studentService *service.StudentService) *StudentController {
	return &StudentController{StudentService: studentService}
}

// Login..
func (con *StudentController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	params := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	in := &service.LoginInput{
		Email:    params.Email,
		Password: params.Password,
	}
	out, err := con.StudentService.Login(ctx, in)
	fmt.Println(err)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		if errors.Is(err, service.ErrInvalidPassword) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}
	}
	token, err := session.CreateToken(c, out.Student)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
