package auth

import (
	"net/http"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func (ac *AuthController) Register(context *gin.Context) {
	var input dto.AuthRequestRegisterDto
	if err := context.ShouldBindJSON(&input); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err := ac.userService.Register(input)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}


	response := dto.SuccessResponse("Success User Registrations", nil)
	context.JSON(http.StatusCreated, response)
}

func (ac *AuthController) Login(context *gin.Context) {
	var input dto.AuthRequestLoginDto
	if err := context.ShouldBindJSON(&input); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := ac.userService.Login(input)

	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := dto.SuccessResponse("Success Login", gin.H{"token": token})
	context.JSON(http.StatusOK, response)
}

