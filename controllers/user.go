package controllers

import (
	"fmt"
	"net/http"

	"github.com/adhityaf/iglobal-be-test/params"
	"github.com/adhityaf/iglobal-be-test/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: *userService,
	}
}

func (u *UserController) Login(ctx *gin.Context) {
	var req params.LoginUser
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

	result := u.userService.Login(req)

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) Register(ctx *gin.Context) {
	var req params.CreateUser
	validate := validator.New()

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	// Request only username, password and name
	// Using register endpoint to make role user
	req.Role = "user"

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

	result := u.userService.CreateUser(req)

	ctx.JSON(result.Status, result.Payload)
}