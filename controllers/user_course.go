package controllers

import (
	"github.com/adhityaf/iglobal-be-test/helpers"
	"github.com/adhityaf/iglobal-be-test/params"
	"github.com/adhityaf/iglobal-be-test/services"
	"github.com/gin-gonic/gin"
)

type UserCourseController struct {
	userCourseService services.UserCourseService
}

func NewUserCourseController(userCourseService *services.UserCourseService) *UserCourseController {
	return &UserCourseController{
		userCourseService: *userCourseService,
	}
}

func (c *UserCourseController) EnrollCourse(ctx *gin.Context) {
	userId := helpers.StringToInt(ctx.MustGet("user_id").(string))
	courseId := helpers.StringToInt(ctx.Param("courseId"))

	var request = params.EnrollCourse{
		UserId:   userId,
		CourseId: courseId,
	}

	result := c.userCourseService.EnrollCourse(request)

	ctx.JSON(result.Status, result.Payload)
}
