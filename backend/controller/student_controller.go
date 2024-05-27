package controller

import (
	"errors"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sichangSun/course-enrollment-manager/domain/model"
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

type StudentCourseControllerOutput struct {
	CoursesList []*model.StudentCoursesWithMulitSchedules
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

	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
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

// ChangePassword
func (con *StudentController) ChangePassword(c echo.Context) error {
	ctx := c.Request().Context()
	claims, err := session.ValidateToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	studentID, err := strconv.Atoi(claims.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	params := struct {
		OldPassword string `json:"oldpassword"`
		NewPassword string `json:"newpassword"`
	}{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	in := &service.UpdatePasswordInput{
		StudentID:   studentID,
		OldPassword: params.OldPassword,
		NewPassword: params.NewPassword,
	}
	err = con.StudentService.UpdatePassword(ctx, in)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}
		if errors.Is(err, service.ErrInvalidPassword) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "successful"})

}

// GetStudentCourses
func (con *StudentController) GetStudentCourses(c echo.Context) error {
	ctx := c.Request().Context()
	claims, err := session.ValidateToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	studentID, err := strconv.Atoi(claims.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	out, err := con.StudentService.GetStudentCourses(ctx, studentID)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return c.JSON(http.StatusOK, map[string]string{"message": "successful but 0 row"})
		}
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	courseList := StudentCourseControllerOutput{out.StudentCourses}
	res := courseList

	return c.JSON(http.StatusOK, res)

}

// RegisterCourse
func (con *StudentController) RegisterCourse(c echo.Context) error {
	ctx := c.Request().Context()
	claims, err := session.ValidateToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	studentID, err := strconv.Atoi(claims.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	params := struct {
		CourseID   string `json:"courseid"`
		CourseName string `json:"coursename"`
	}{}
	if err := c.Bind(&params); err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	couserID, err := strconv.Atoi(params.CourseID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	in := service.RegisterCourseInput{
		StudentID: studentID,
		CourseID:  couserID,
	}

	err = con.StudentService.RegisterCourse(ctx, &in)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "successful"})
}

// UnregisterCourse
func (con *StudentController) UnRegisterCourse(c echo.Context) error {
	ctx := c.Request().Context()
	claims, err := session.ValidateToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	studentID, err := strconv.Atoi(claims.ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	ID := c.Param("course_id")
	courseID, err := strconv.Atoi(ID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	in := &service.UnRegisterCourseInput{
		StudentID: studentID,
		CourseID:  courseID,
	}
	err = con.StudentService.UnRegisterCourse(ctx, in)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "successful"})
}
