package main

import (
	"log"

	"github.com/adhityaf/iglobal-be-test/config"
	"github.com/adhityaf/iglobal-be-test/controllers"
	"github.com/adhityaf/iglobal-be-test/middlewares"
	"github.com/adhityaf/iglobal-be-test/repositories"
	"github.com/adhityaf/iglobal-be-test/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	route := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	courseRepository := repositories.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepository)
	courseController := controllers.NewCourseController(courseService)

	userCourseRepository := repositories.NewUserCourseRepository(db)
	userCourseService := services.NewUserCourseService(userCourseRepository, courseRepository)
	userCourseController := controllers.NewUserCourseController(userCourseService)

	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/login", userController.Login)
		mainRouter.POST("/register", userController.Register)

		authorized := mainRouter.Group("/")
		authorized.Use(middlewares.Auth())
		{
			authorized.GET("/courses/search", courseController.GetCourseBySearch)
			authorized.GET("/courses/:courseId", courseController.GetCourseById)

			user := authorized.Group("/")
			user.Use(middlewares.IsUser())
			{
				authorized.POST("/courses/enroll/:courseId", userCourseController.EnrollCourse)
			}

			admin := authorized.Group("/")
			admin.Use(middlewares.IsAdmin())
			{
				admin.POST("/courses", courseController.CreateCourse)
				admin.GET("/courses", courseController.GetAllCourses)
				admin.PUT("/courses/:courseId", courseController.UpdateCourseById)
				admin.DELETE("/courses/:courseId", courseController.DeleteCourseById)
			}
		}
	}

	route.Run()
}
