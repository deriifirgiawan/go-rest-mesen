package services

import (
	"fmt"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/models"
	"rest-app-pos/src/repository"
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
	merchantRepository := repository.NewMerchantRepository()

	merchant, _ := merchantRepository.FindByUserId(user_id)
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
	merchantRepository := repository.NewMerchantRepository()

	merchant, _ := merchantRepository.FindByUserId(user_id)
	product := &models.Product{
		Name: payload.Name,
		Description: payload.Description,
		MerchantID: merchant.ID,
		CategoryID: uint(payload.CategoryID),
		Price: num,
	}

	return s.repo.Update(product)
}


func (s *productService) DeleteProduct(id uint, user_id uint) error {
	merchantRepository := repository.NewMerchantRepository()

	merchant, _ := merchantRepository.FindByUserId(user_id)
	product, err := s.repo.FindProductById(id)
	if err != nil {
		return err
	}

	if product.MerchantID != merchant.ID {
		return fmt.Errorf("unauthorized: you don't have permission to delete this product")
	}

	return s.repo.Delete(id)
}