package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sichangSun/course-enrollment-manager/controller"
	"github.com/sichangSun/course-enrollment-manager/domain/service"
	"github.com/sichangSun/course-enrollment-manager/infrastructure/mysql"
)

// func NewRouter() *Router {
// 	e := echo.New()
// 	e.Use(middleware.Recover())
// 	e.Use(middleware.CSRF())
// 	e.Use(middleware.JWT())
// }
//
// login
// password変更

// 登録した授業一覧
// 全授業返す
// 授業詳細返す
// 授業登録
// 授業修正
// 授業削除

// Router ...
type Router struct {
	e *echo.Echo
}

// RouterConfig ...
type RouterConfig struct {
	DB *sqlx.DB
}

// New ...
func New(conf *RouterConfig) *Router {
	e := echo.New()

	e.Use(middleware.Recover())
	// e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	courseRepository := mysql.NewCourseRepository(conf.DB)
	courseService := service.NewCourseService(courseRepository)
	courseController := controller.NewCourseController(courseService)

	{
		g := e.Group("/api")
		{
			g3 := g.Group("/courses")
			g3.GET("", courseController.GetAllCourses)
			g3.GET("/:id", courseController.GetOneCourseDetail)
		}
	}

	return &Router{e}
}

// Run ...
func (r *Router) Run(addr string) error {
	return r.e.Start(addr)
}
