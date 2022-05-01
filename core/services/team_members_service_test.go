package services

import (
	"github.com/stretchr/testify/require"
	"merchant-account-api/core/domain"
	"testing"
)

func TestValidateTeamMember(t *testing.T) {
	err := validateTeamMember("save", domain.TeamMember{
		Name:  "abc",
		Email: "def@gmail.com",
	})
	require.NoError(t, err)
}

func TestValidateTeamMemberInvalidEmail(t *testing.T) {
	err := validateTeamMember("save", domain.TeamMember{
		Name:  "abc",
		Email: "def",
	})
	require.Error(t, err)
}
