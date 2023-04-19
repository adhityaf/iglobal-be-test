package services

import (
	"net/http"

	"github.com/adhityaf/iglobal-be-test/models"
	"github.com/adhityaf/iglobal-be-test/params"
	"github.com/adhityaf/iglobal-be-test/repositories"
	"github.com/gin-gonic/gin"
)

type UserCourseService struct {
	userCourseRepository repositories.UserCourseRepository
	courseRepository repositories.CourseRepository
}

func NewUserCourseService(userCourseRepository repositories.UserCourseRepository, courseRepository repositories.CourseRepository) *UserCourseService {
	return &UserCourseService{
		userCourseRepository: userCourseRepository,
		courseRepository: courseRepository,
	}
}

func (u *UserCourseService) EnrollCourse(request params.EnrollCourse) *params.Response {
	_, err := u.courseRepository.FindById(request.CourseId)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			},
		}
	}

	_, err = u.userCourseRepository.FindById(request.CourseId)
	if err == nil {
		return &params.Response{
			Status: http.StatusOK,
			Payload: gin.H{
				"message": "You already enroll this course",
			},
		}
	}

	userCourse := models.UserCourse{
		UserId:   request.UserId,
		CourseId: request.CourseId,
	}

	courseData, err := u.userCourseRepository.Create(&userCourse)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusCreated,
		Payload: gin.H{
			"data": courseData,
		},
	}
}
