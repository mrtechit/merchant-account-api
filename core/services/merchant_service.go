package services

import (
	"errors"
	"github.com/badoux/checkmail"
	"merchant-account-api/core/domain"
	"merchant-account-api/ports/api"
	"merchant-account-api/ports/repository"
	"strings"
)

type MerchantService struct {
	merchantRepo repository.MerchantRepository
}

func NewMerchantService(merchantRepo repository.MerchantRepository) api.MerchantServiceInterface {
	return &MerchantService{merchantRepo: merchantRepo}
}

func (m MerchantService) Save(merchant domain.Merchant) (string, error) {
	err := validateMerchant("save", merchant)
	if err != nil {
		return "", err
	}
	return m.merchantRepo.SaveMerchant(merchant)
}

func (m MerchantService) Update(merchant domain.Merchant) error {
	err := validateMerchant("", merchant)
	if err != nil {
		return err
	}
	return m.merchantRepo.UpdateMerchant(merchant)
}

func (m MerchantService) Delete(id string) error {
	return m.merchantRepo.DeleteMerchant(id)
}

func (m MerchantService) Find(id string) (domain.Merchant, error) {
	return m.merchantRepo.FindMerchantByID(id)
}

func validateMerchant(action string, merchant domain.Merchant) error {
	switch strings.ToLower(action) {
	case "save":
		if merchant.MerchantName == "" {
			return errors.New("Required Nickname")
		}
		if merchant.MerchantCity == "" {
			return errors.New("Required MerchantCity")
		}
		if merchant.MerchantCategoryCode == "" {
			return errors.New("Required MerchantCategoryCode")
		}
		if merchant.MerchantEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(merchant.MerchantEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if merchant.MerchantCode == 0 {
			return errors.New("Required Merchant Code")
		}
		if merchant.MerchantEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(merchant.MerchantEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
