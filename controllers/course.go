package controllers

import (
	"fmt"
	"net/http"

	"github.com/adhityaf/iglobal-be-test/helpers"
	"github.com/adhityaf/iglobal-be-test/params"
	"github.com/adhityaf/iglobal-be-test/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CourseController struct {
	courseService services.CourseService
}

func NewCourseController(courseService *services.CourseService) *CourseController {
	return &CourseController{
		courseService: *courseService,
	}
}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var req params.CreateCourse
	validate := validator.New()

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	err = validate.Struct(req)
	if err != nil {
		validationMessage := ""
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage = fmt.Sprintf("%s field %s %s. ", validationMessage, err.Field(), err.Tag())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: validationMessage,
		})

		return
	}

	result := c.courseService.CreateCourse(req)

	ctx.JSON(result.Status, result.Payload)
}

func (c *CourseController) GetAllCourses(ctx *gin.Context) {
	result := c.courseService.GetAllCourses()

	ctx.JSON(result.Status, result.Payload)
}

func (c *CourseController) GetCourseById(ctx *gin.Context) {
	courseId := helpers.StringToInt(ctx.Param("courseId"))

	result := c.courseService.GetCourseById(courseId)

	ctx.JSON(result.Status, result.Payload)
}

func (c *CourseController) UpdateCourseById(ctx *gin.Context) {
	var req params.UpdateCourse
	validate := validator.New()

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	// Get user_id from param
	courseId := helpers.StringToInt(ctx.Param("courseId"))

	req.CourseId = courseId

	err = validate.Struct(req)
	if err != nil {
		validationMessage := ""
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage = fmt.Sprintf("%s field %s %s. ", validationMessage, err.Field(), err.Tag())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: validationMessage,
		})

		return
	}

	result := c.courseService.UpdateCourse(req)

	ctx.JSON(result.Status, result.Payload)
}

func (c *CourseController) DeleteCourseById(ctx *gin.Context) {
	// Get course_id from param
	courseId := helpers.StringToInt(ctx.Param("courseId"))

	result := c.courseService.DeleteCourse(courseId)

	ctx.JSON(result.Status, result.Payload)
}

func (c *CourseController) GetCourseBySearch(ctx *gin.Context) {
	query := ctx.Query("query")
	page := helpers.StringToInt(ctx.Query("page"))
	size := helpers.StringToInt(ctx.Query("size"))
	level := ctx.Query("level")
	language := ctx.Query("language")

	criteria := params.CourseSearchCriteria{
		Query:    query,
		Page:     page,
		Size:     size,
		Level:    level,
		Language: language,
	}

	result := c.courseService.GetAllCourseBySearch(criteria)

	ctx.JSON(result.Status, result.Payload)
}
