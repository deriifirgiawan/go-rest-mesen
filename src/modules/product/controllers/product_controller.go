package product

import (
	"rest-app-pos/src/dto"
	"rest-app-pos/src/modules/product/services"
	"rest-app-pos/src/utils"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) GetAllProducts(context *gin.Context) {
	var filter dto.ProductQueryFilterDto

	if err := context.ShouldBindQuery(&filter); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	products := pc.productService.GetAllProducts(filter)

	if products.Error != "" {
		response := dto.ErrorResponse(http.StatusBadRequest, products.Error)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := dto.SuccessResponseWithPagination(
		"Success Get List Product",
		*products.List,
		*products.Limit,
		*products.Page,
		int(*products.TotalItems),
		*products.TotalPages,
	)
	context.JSON(http.StatusOK, response)
}

func (pc *ProductController) CreateProduct(context *gin.Context) {
	var input dto.ProductRequestDto
	decodeToken := utils.DecodeToken(context)

	if err := context.ShouldBindJSON(&input); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err := pc.productService.InsertProduct(input, decodeToken.User_ID)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	response := dto.SuccessResponse("Success Add Product", input)
	context.JSON(http.StatusCreated, response)
}

func (pc *ProductController) UpdateProduct(context *gin.Context) {
	var input dto.ProductRequestUpdateDto
	decodeToken := utils.DecodeToken(context)

	if err := context.ShouldBindJSON(&input); err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err := pc.productService.UpdateProduct(input, decodeToken.User_ID)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	response := dto.SuccessResponse("Success Add Product", input)
	context.JSON(http.StatusCreated, response)
}

func (pc *ProductController) DeleteProduct(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, "Invalid product ID")
		context.JSON(http.StatusBadRequest, response)
		return
	}

	decodeToken := utils.DecodeToken(context)

	err = pc.productService.DeleteProduct(uint(id), decodeToken.User_ID)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := dto.SuccessResponse("Product deleted successfully", nil)
	context.JSON(http.StatusOK, response)
}

func (pc *ProductController) GetAllCategories(context *gin.Context) {
	categories, err := pc.productService.GetAllCategories()

	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := dto.SuccessResponse("Success Get All Categories", categories)
	context.JSON(http.StatusOK, response)
}

func (pc *ProductController) GetCategoryById(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 10)
	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, "Invalid product ID")
		context.JSON(http.StatusBadRequest, response)
		return
	}
	category, err := pc.productService.GetCategoryById(uint(id))

	if err != nil {
		response := dto.ErrorResponse(http.StatusBadRequest, err.Error())
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response := dto.SuccessResponse("Success Get All Category", category)
	context.JSON(http.StatusOK, response)
}