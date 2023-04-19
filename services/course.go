package services

import (
	"fmt"
	"net/http"

	"github.com/adhityaf/iglobal-be-test/models"
	"github.com/adhityaf/iglobal-be-test/params"
	"github.com/adhityaf/iglobal-be-test/repositories"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type CourseService struct {
	courseRepository repositories.CourseRepository
}

func NewCourseService(repo repositories.CourseRepository) *CourseService {
	return &CourseService{
		courseRepository: repo,
	}
}

func (c *CourseService) CreateCourse(request params.CreateCourse) *params.Response {
	course := models.Course{
		Title:       request.Title,
		Slug:        slug.Make(request.Title),
		Description: request.Description,
		Price:       request.Price,
		Level:       request.Level,
		Language:    request.Language,
	}

	courseData, err := c.courseRepository.Create(&course)
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

func (c *CourseService) UpdateCourse(request params.UpdateCourse) *params.Response {
	course, err := c.courseRepository.FindById(request.CourseId)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			},
		}
	}

	course.Title = request.Title             // update title
	course.Slug = slug.Make(request.Title)   // update slug
	course.Description = request.Description // update description
	course.Price = request.Price             // update price
	course.Level = request.Level             // update level
	course.Language = request.Language       // update language

	course, err = c.courseRepository.Update(course)
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
		Status: http.StatusOK,
		Payload: gin.H{
			"message": fmt.Sprintf("Success update data with id : %d", course.CourseId),
		},
	}
}

func (c *CourseService) DeleteCourse(courseId int) *params.Response {
	course, err := c.courseRepository.FindById(courseId)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			},
		}
	}

	course, err = c.courseRepository.Delete(course)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": fmt.Sprintf("Success delete data with id : %d", course.CourseId),
		},
	}
}

func (c *CourseService) GetCourseById(courseId int) *params.Response {
	course, err := c.courseRepository.FindById(courseId)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"data": course,
		},
	}
}

func (c *CourseService) GetAllCourses() *params.Response {
	courses, err := c.courseRepository.FindAllCourses()
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"items": courses,
		},
	}
}

func (c *CourseService) GetAllCourseBySearch(criteria params.CourseSearchCriteria) *params.Response {
	criteria.SetDefaultValue()
	courses, err := c.courseRepository.FindAllBySearchAndFilter(criteria)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"items": courses,
		},
	}
}
