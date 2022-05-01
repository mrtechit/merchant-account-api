package middleware

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestGenerateNewAccessTokenSuccess(t *testing.T) {
	initEnv()
	_, err := generateNewAccessToken()
	require.NoError(t, err)

}

func initEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}
