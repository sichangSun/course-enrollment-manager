package controller

import (
	"errors"
	"net/http"
	"strconv"

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
	CourseList []*model.CourseWithMulitSchedules
}
type OneCourseControllerOutput struct {
	CourseDetail *model.CourseWithMulitSchedules
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
	res := courseList

	return c.JSON(http.StatusOK, res)
}

// GetGetOneCourseDetail ...
func (con *CourseController) GetOneCourseDetail(c echo.Context) error {
	ctx := c.Request().Context()
	courseID := c.Param("course_id")

	id, err := strconv.Atoi(courseID)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	out, err := con.CourseService.GetGetOneCourseDetail(ctx, id)
	if err != nil {
		c.Logger().Error(err.Error())
		if errors.Is(err, service.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	res := OneCourseControllerOutput{out.Courses}
	return c.JSON(http.StatusOK, res)

}
