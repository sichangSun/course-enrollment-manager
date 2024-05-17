package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sichangSun/course-enrollment-manager/domain/model"
	"github.com/sichangSun/course-enrollment-manager/domain/service"
)

// CourseController ...
type CourseController struct {
	CourseService *service.CourseService
}

type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Semester    int    `json:"semester"`
	Instructor  int    `json:"instructor"`
	Credits     int    `json:"credits"`
	Description string `json:"description"`
}
type CourseControllerOutput struct {
	CourseService []*model.CourseWithMulitSchedules
}

func NewCourseController(cs *service.CourseService) *CourseController {
	return &CourseController{
		CourseService: cs,
	}
}

// GetAllCourses ...
func (con *CourseController) GetAllCourses(c echo.Context) error {
	ctx := c.Request().Context()

	out, err := con.CourseService.GetAllCourses(ctx)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	courseList := CourseControllerOutput{out.Courses}
	res := struct {
		CourseList CourseControllerOutput `json:"course_list"`
	}{
		CourseList: courseList,
	}

	return c.JSON(http.StatusOK, res)
}
