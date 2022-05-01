//go:build integration
// +build integration

package sql_repository

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"log"
	"merchant-account-api/core/domain"
	"testing"
)

func TestSaveTeamMember_UniqueKeyViolated_Failed(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	teamMemberRepo, _ := NewPostgreSQLRepository()
	err := teamMemberRepo.SaveTeamMember(domain.TeamMember{
		Name:                 "tmp",
		Email:                "tmp2@gmail.com",
		MerchantMerchantCode: 2,
	})

	err = teamMemberRepo.SaveTeamMember(domain.TeamMember{
		Name:  "tmp",
		Email: "tmp2@gmail.com",
	})
	require.Error(t, err)
}

func TestUpdateTeamMember(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	teamMemberRepo, _ := NewPostgreSQLRepository()
	err := teamMemberRepo.SaveTeamMember(domain.TeamMember{
		Name:                 "tmp",
		Email:                "tmp1@gmail.com",
		MerchantMerchantCode: 1,
	})
	require.NoError(t, err)
}

func TestSaveTeamMember(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	teamMemberRepo, _ := NewPostgreSQLRepository()
	err := teamMemberRepo.SaveTeamMember(domain.TeamMember{
		Name:                 "tmp",
		Email:                "tmp4@gmail.com",
		MerchantMerchantCode: 2,
	})
	require.NoError(t, err)
}

func TestDeleteTeamMember(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	teamMemberRepo, _ := NewPostgreSQLRepository()
	err := teamMemberRepo.DeleteTeamMember(2)
	require.Error(t, err)
}

func TestGetTeamMembertById(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	teamMemberRepo, _ := NewPostgreSQLRepository()
	result, err := teamMemberRepo.FindTeamMemberByID(2)
	require.NoError(t, err)
	require.Equal(t, "tmp2@gmail.com", result.Email)
}

func TestFindTeamMembersByMerchantCode(t *testing.T) {
	//startEmbeddedDB()
	initEnv()
	teamMemberRepo, _ := NewPostgreSQLRepository()
	result, err := teamMemberRepo.FindTeamMembersByMerchantCode("2", 2, 5)
	fmt.Println(result)
	require.NoError(t, err)
}

func initEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}

/*func startEmbeddedDB() *embeddedpostgres.EmbeddedPostgres {

	postgres := embeddedpostgres.NewDatabase()
	err := postgres.Start()
	if err != nil {
		fmt.Println("test", err)
	}
	return postgres
}
*/
