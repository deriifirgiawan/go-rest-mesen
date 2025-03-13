package services

import (
	"errors"
	"fmt"
	"rest-app-pos/src/dto"
	"rest-app-pos/src/models"
	"rest-app-pos/src/modules/merchant/repository"
)

type MerchantService interface {
	GetMerchantByUserId(user_id uint) (*models.Merchant, error)
	CreateMerchant(payload dto.MerchantRequestDto, user_id uint) error
	UpdateMerchant(payload dto.MerchantRequestDto, user_id uint)error
}

type merchantService struct {
	repo repository.MerchantRepository
}

func NewMerchantService(repo repository.MerchantRepository) MerchantService {
	return &merchantService{repo: repo}
}

func (s *merchantService) GetMerchantByUserId(user_id uint) (*models.Merchant, error) {
	merchant, err := s.repo.FindByUserId(user_id)
	if err != nil {
		return nil, errors.New("merchant not found")
	}
	return merchant, nil
}

func (s *merchantService) CreateMerchant(payload dto.MerchantRequestDto, user_id uint) error {
	existingMerchant, _ := s.repo.FindByUserId(user_id)


	if existingMerchant != nil {
		return errors.New("you already have a merchant")
	}

	merchant := &models.Merchant{
		Name:   payload.Name,
		UserID: user_id,
	}

	if err := s.repo.Create(merchant); err != nil {
		return fmt.Errorf("failed to create merchant: %v", err)
	}

	return nil
}


func (s *merchantService) UpdateMerchant(payload dto.MerchantRequestDto, user_id uint) error {
	merchant := &models.Merchant{
		Name: payload.Name,
	}

	return s.repo.Update(merchant, user_id)
}