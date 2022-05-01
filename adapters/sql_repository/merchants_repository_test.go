//go:build integration
// +build integration

package sql_repository

import (
	"github.com/stretchr/testify/require"
	"merchant-account-api/core/domain"
	"testing"
)

func TestGetMerchantByIdFailed(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	merchantRepo, _ := NewPostgreSQLRepository()
	_, err := merchantRepo.FindMerchantByID("1")
	require.Error(t, err)
	require.EqualError(t, err, "record not found")
}

func TestSaveMerchant(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	merchantRepo, _ := NewPostgreSQLRepository()
	err := merchantRepo.SaveMerchant(domain.Merchant{
		MerchantName:         "visa",
		MerchantEmail:        "visa@gmail.com",
		MerchantCity:         "bkk",
		MerchantCategoryCode: "1234",
	})
	require.NoError(t, err)
}

func TestUpdateMerchant(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	merchantRepo, _ := NewPostgreSQLRepository()
	err := merchantRepo.UpdateMerchant(domain.Merchant{
		MerchantCode: 1,
		MerchantCity: "thailand",
	})
	require.NoError(t, err)
}

func TestDeleteMerchant(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	merchantRepo, _ := NewPostgreSQLRepository()
	err := merchantRepo.DeleteMerchant("1")
	require.NoError(t, err)
}
