package services

import (
	"errors"
	"fmt"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/models"
	merchant "rest-app-pos/src/modules/merchant/repository"
	"rest-app-pos/src/modules/product/repository"
	"strconv"
)


type ResponseGetAllProduct struct {
	Products []models.ProductList
}

type ProductService interface {
	InsertProduct(payload dto.ProductRequestDto, user_id uint) error
	UpdateProduct(payload dto.ProductRequestUpdateDto, user_id uint) error
	DeleteProduct(id uint, user_id uint) error
	GetAllProducts(filter dto.ProductQueryFilterDto) dto.ResponseGetListProduct
	GetAllProductsByOwner(user_id uint) ([]models.ProductList, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryById(id uint) (*models.Category, error)
}

type productService struct {
	repo repository.ProductRepository
	repoCategory repository.CategoryRepository
}

func NewProductService(repo repository.ProductRepository, repoCategory repository.CategoryRepository) ProductService {
	return &productService{repo: repo, repoCategory: repoCategory}
}

func (s *productService) InsertProduct(payload dto.ProductRequestDto, user_id uint) error {
	num, err := strconv.ParseFloat(payload.Price, 64)
	if err != nil {
		return fmt.Errorf("invalid price format: %v", err)
	}

	merchant, err := merchant.NewMerchantRepository().FindByUserId(user_id)
	if err != nil {
		return fmt.Errorf("merchant not found for user ID %d: %v", user_id, err)
	}

	if merchant == nil {
		return errors.New("merchant not found")
	}

	product := &models.Product{
		Name:        payload.Name,
		Description: payload.Description,
		MerchantID:  merchant.ID,
		CategoryID:  uint(payload.CategoryID),
		Price:       num,
		Quantity: payload.Quantity,
	}

	return s.repo.Create(product)
}


func (s *productService) UpdateProduct(payload dto.ProductRequestUpdateDto, user_id uint) error {
	num, err := strconv.ParseFloat(payload.Price, 64)
	if err != nil {
		return fmt.Errorf("invalid price format: %v", err)
	}

	merchant, err := merchant.NewMerchantRepository().FindById(user_id)
	if err != nil {
		return fmt.Errorf("merchant not found for user ID %d: %v", user_id, err)
	}

	if merchant == nil {
		return errors.New("merchant not found")
	}

	product := &models.Product{
		ID: payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		MerchantID: merchant.ID,
		CategoryID:  uint(payload.CategoryID),
		Price:       num,
		Quantity: payload.Quantity,
	}

	return s.repo.Update(product, merchant.ID)
}


func (s *productService) DeleteProduct(id uint, user_id uint) error {
	merchantRepository := merchant.NewMerchantRepository()

	merchants, _ := merchantRepository.FindByUserId(user_id)
	product, err := s.repo.FindProductById(id)
	if err != nil {
		return err
	}

	if product.MerchantID != merchants.ID {
		return fmt.Errorf("unauthorized: you don't have permission to delete this product")
	}

	return s.repo.Delete(id)
}

func (s *productService) GetAllProducts(filter dto.ProductQueryFilterDto) dto.ResponseGetListProduct{
	products, totalItems, err := s.repo.GetAllProductByMerchant(filter)
	if err != nil {
		return dto.ResponseGetListProduct{
			Error: err.Error(),
		}
	}

	limit := 10

	if filter.Limit != nil && *filter.Limit > 0 {
		limit = *filter.Limit
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))
	return dto.ResponseGetListProduct{
		List: &products,
		Page: filter.Page,
		Limit: &limit,
		TotalItems: &totalItems,
		TotalPages: &totalPages,
		Error: "",
	}
}

func (s *productService) GetAllProductsByOwner(user_id uint) ([]models.ProductList, error) {
	merchant, err := merchant.NewMerchantRepository().FindById(user_id)
	if err != nil {
		return nil, fmt.Errorf("merchant not found for user ID %d: %v", user_id, err)
	}

	if merchant == nil {
		return nil, errors.New("merchant not found")
	}

	filter := &dto.ProductQueryFilterDto{
		MerchantID: &merchant.ID,
	}
	products, _, err := s.repo.GetAllProductByMerchant(*filter)
	if err != nil {
		return nil, errors.New("merchant not found")
	}
	return products, nil
}

func (s *productService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.repoCategory.GetAllCategories()
	
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return categories, nil
}

func (s *productService) GetCategoryById(id uint) (*models.Category, error) {
	category, err := s.repoCategory.GetCategoryById(id)
	
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return category, nil
}