package merchant

import (
	"net/http"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/modules/merchant/services"
	"rest-app-pos/src/utils"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantService services.MerchantService
}

func NewMerchantController(merchantService services.MerchantService) *MerchantController {
	return &MerchantController{merchantService: merchantService}
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