package services

import (
	"fmt"
	"net/http"

	"github.com/adhityaf/iglobal-be-test/helpers"
	"github.com/adhityaf/iglobal-be-test/models"
	"github.com/adhityaf/iglobal-be-test/params"
	"github.com/adhityaf/iglobal-be-test/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (u *UserService) Login(request params.LoginUser) *params.Response {
	user, err := u.userRepository.FindByEmail(request.Email)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("User %s", err.Error()),
			},
		}
	}

	ok := helpers.ComparePassword([]byte(user.Password), []byte(request.Password))
	if !ok {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"error":   "Unauthorized",
				"message": "Password not match",
			},
		}
	}

	token := helpers.GenerateToken(user.UserId, user.Name, user.Role)

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "Login successful",
			"token":   token,
		},
	}
}

func (u *UserService) CreateUser(request params.CreateUser) *params.Response {
	password := helpers.HashPassword(request.Password)
	user := models.User{
		Email:    request.Email,
		Password: password,
		Name:     request.Name,
		Role:     request.Role,
	}

	userData, err := u.userRepository.Create(&user)
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
			"data": userData,
		},
	}
}