package sql_repository

import (
	"fmt"
	"merchant-account-api/core/domain"
	"strconv"
)

func (p *postgreSQLRepository) FindMerchantByID(id string) (domain.Merchant, error) {
	merchant := domain.Merchant{}
	result := p.db.First(&merchant, "merchant_code = ?", id)
	if result.Error != nil {
		return merchant, result.Error
	}

	return merchant, nil
}

func (p *postgreSQLRepository) SaveMerchant(merchant domain.Merchant) (string, error) {

	merchantSave := domain.Merchant{
		MerchantName:         merchant.MerchantName,
		MerchantEmail:        merchant.MerchantEmail,
		MerchantCity:         merchant.MerchantCity,
		MerchantCategoryCode: merchant.MerchantCategoryCode,
	}
	merchantResult := p.db.Create(&merchantSave)
	if merchantResult.Error != nil {
		return "", merchantResult.Error
	}
	fmt.Println("save merchant success")
	return strconv.Itoa(int(merchantSave.MerchantCode)), nil

}

func (p *postgreSQLRepository) UpdateMerchant(merchant domain.Merchant) error {
	merchantUpdate := domain.Merchant{}
	result := p.db.Model(&merchantUpdate).Where("merchant_code = ?", merchant.MerchantCode).Updates(domain.Merchant{
		MerchantName:         merchant.MerchantName,
		MerchantCity:         merchant.MerchantCity,
		MerchantEmail:        merchant.MerchantEmail,
		MerchantCategoryCode: merchant.MerchantCategoryCode,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *postgreSQLRepository) DeleteMerchant(merchantCode string) error {
	result := p.db.Where("merchant_code = ?", merchantCode).Delete(&domain.Merchant{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
