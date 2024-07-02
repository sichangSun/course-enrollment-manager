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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders:     []string{"Content-Type", "X-Csrf-Token", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"X-CSRF-Token"},
	}))
	// e.Use(middleware.CSRF())

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
			g1 := g.Group("")
			g1.POST("/login", studentController.Login)
		}
		{
			g2 := g.Group("")
			g2.Use(echojwt.WithConfig(config))
			g2.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
				CookieSecure: false, //localhost
				CookiePath:   "/api",
			}))
			{
				g3 := g2.Group("/auth")
				g3.PUT("/change-password", studentController.ChangePassword)
				g3.GET("/courses", studentController.GetStudentCourses)
				g3.POST("/course", studentController.RegisterCourse)
				g3.DELETE("/course/:course_id", studentController.UnRegisterCourse)

			}
			{
				g4 := g2.Group("/courses")
				g4.GET("", courseController.GetAllCourses)
				g4.GET("/:course_id", courseController.GetOneCourseDetail)
			}

		}

	}

	return &Router{e}
}

// Run ...
func (r *Router) Run(addr string) error {
	return r.e.Start(addr)
}
