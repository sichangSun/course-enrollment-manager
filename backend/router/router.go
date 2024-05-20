package router

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sichangSun/course-enrollment-manager/controller"
	"github.com/sichangSun/course-enrollment-manager/domain/service"
	"github.com/sichangSun/course-enrollment-manager/infrastructure/mysql"
	"github.com/sichangSun/course-enrollment-manager/session"
)

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

	//student init
	studentRepository := mysql.NewStudentRepository(conf.DB)
	studentService := service.NewStudentService(studentRepository)
	studentController := controller.NewStudentController(studentService)

	// JWT middleware setting
	config := echojwt.Config{
		SigningKey: []byte("secret"),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(session.AccountClaims)
		},
	}
	{
		g := e.Group("/api")
		{
			g1 := g.Group("/auth")
			g1.Use(echojwt.WithConfig(config))
			g1.PUT("/change-password", studentController.ChangePassword)
			g1.GET("/courses", studentController.GetStudentCourses)
			g1.POST("/course", studentController.RegisterCourse)

		}
		{
			g2 := g.Group("/auth")
			g2.POST("/login", studentController.Login)
		}
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
