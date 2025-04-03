package merchant

import (
	"net/http"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/modules/merchant/services"
	userService "rest-app-pos/src/services"
	"rest-app-pos/src/utils"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantService services.MerchantService
	userService userService.UserService
}

func NewMerchantController(merchantService services.MerchantService, userService userService.UserService) *MerchantController {
	return &MerchantController{merchantService: merchantService, userService: userService}
}

func (mc *MerchantController) AddMerchant(context *gin.Context) {
	var input dto.MerchantRequestDto

	decodeToken := utils.DecodeToken(context)

	if err := context.ShouldBindJSON(&input); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err := mc.merchantService.CreateMerchant(input, decodeToken.User_ID)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	response := dto.SuccessResponse("Success Create New Merchant", input)
	context.JSON(http.StatusCreated, response)
}

func (mc *MerchantController) UpdateMerchant(context *gin.Context) {
	var input dto.MerchantRequestDto

	decodeToken := utils.DecodeToken(context)

	if err := context.ShouldBindJSON(&input); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err := mc.merchantService.UpdateMerchant(input, decodeToken.User_ID)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	response := dto.SuccessResponse("Success Create New Merchant", input)
	context.JSON(http.StatusCreated, response)
}

func (mc *MerchantController) GetMerchant(context *gin.Context) {
	decodeToken := utils.DecodeToken(context).User_ID
	merchant, err := mc.merchantService.GetMerchantByUserId(decodeToken)

	if err != nil {
		response := dto.ErrorResponse(http.StatusNotFound, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	response := dto.SuccessResponse("Success Get Merchant", merchant)
	context.JSON(http.StatusOK, response)
}

func (mc *MerchantController) AddEmployee(context *gin.Context) {
	var input dto.MerchantUserRequestDto

	if err := context.ShouldBindJSON(&input); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := mc.userService.AddEmployee(input)

	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	merchantUser := &dto.MerchantUserRequestDto{
		MerchantID: input.MerchantID,
		UserID: user.ID,
	}

	err = mc.merchantService.AddNewEmployee(*merchantUser)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	response := dto.SuccessResponse("Success Add New Employee", input)
	context.JSON(http.StatusCreated, response)
}

func (mc *MerchantController) GetAllEmployee(context *gin.Context) {
	decodeToken := utils.DecodeToken(context)
	merchant, err := mc.merchantService.GetMerchantByUserId(decodeToken.User_ID)
	if err != nil {
		response := dto.ErrorResponse(http.StatusNotFound, err.Error())
		context.JSON(http.StatusNotFound, response)

		return
	}

	employee, err := mc.merchantService.GetAllEmployee(merchant.ID)

	if err != nil {
		response := dto.ErrorResponse(http.StatusNotFound, err.Error())
		context.JSON(http.StatusNotFound, response)

		return
	}

	response := dto.SuccessResponse("Success Get Employees", employee)
	context.JSON(http.StatusOK, response)
}