package services

import (
	"fmt"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/models"
	merchant "rest-app-pos/src/modules/merchant/repository"
	"rest-app-pos/src/modules/product/repository"
	"strconv"
)

type ProductService interface {
	InsertProduct(payload dto.ProductRequestDto, user_id uint) error
	UpdateProduct(payload dto.ProductRequestDto, user_id uint) error
	DeleteProduct(id uint, user_id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) InsertProduct(payload dto.ProductRequestDto, user_id uint) error {
	num, _ := strconv.ParseFloat(payload.Price, 64)
	merchantRepository := repository.NewProductRepository()

	merchant, _ := merchantRepository.FindProductById(user_id)
	product := &models.Product{
		Name: payload.Name,
		Description: payload.Description,
		MerchantID: merchant.ID,
		CategoryID: uint(payload.CategoryID),
		Price: num,
	}

	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(payload dto.ProductRequestDto, user_id uint) error {
	_, err := s.repo.FindProductById(user_id);

	if err != nil {
		return err
	}
	num, _ := strconv.ParseFloat(payload.Price, 64)
	merchantRepository := merchant.NewMerchantRepository()

	merchants, _ := merchantRepository.FindByUserId(user_id)
	product := &models.Product{
		Name: payload.Name,
		Description: payload.Description,
		MerchantID: merchants.ID,
		CategoryID: uint(payload.CategoryID),
		Price: num,
	}

	return s.repo.Update(product)
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