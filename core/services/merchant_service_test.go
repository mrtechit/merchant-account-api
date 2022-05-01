package services

import (
	"github.com/stretchr/testify/require"
	"merchant-account-api/core/domain"
	"testing"
)

func TestValidateMerchant(t *testing.T) {

	err := validateMerchant("save", domain.Merchant{
		MerchantCode:         1,
		MerchantName:         "abc",
		MerchantEmail:        "def@gmail.com",
		MerchantCity:         "xyz",
		MerchantCategoryCode: "123",
	})

	require.NoError(t, err)

}

func TestValidateMerchantInvalidEmail(t *testing.T) {

	err := validateMerchant("save", domain.Merchant{
		MerchantCode:         1,
		MerchantName:         "abc",
		MerchantEmail:        "def",
		MerchantCity:         "xyz",
		MerchantCategoryCode: "123",
	})

	require.Error(t, err)

}
