package api

import (
	"merchant-account-api/core/domain"
)

type MerchantServiceInterface interface {
	Find(id string) (domain.Merchant, error)
	Save(merchant domain.Merchant) (string, error)
	Update(merchant domain.Merchant) error
	Delete(id string) error
}
