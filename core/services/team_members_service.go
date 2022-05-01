package services

import (
	"errors"
	"github.com/badoux/checkmail"
	"merchant-account-api/core/domain"
	"merchant-account-api/ports/api"
	"merchant-account-api/ports/repository"
	"strconv"
	"strings"
)

type TeamMembersService struct {
	teamMembersRepo repository.TeamMembersRepository
}

func NewTeamMembersService(teamMembersRepo repository.TeamMembersRepository) api.TeamMembersServiceInterface {
	return &TeamMembersService{teamMembersRepo: teamMembersRepo}
}

func (t TeamMembersService) Save(teamMember domain.TeamMember) error {
	err := validateTeamMember("save", teamMember)
	if err != nil {
		return err
	}
	return t.teamMembersRepo.SaveTeamMember(teamMember)
}

func (t TeamMembersService) Update(teamMember domain.TeamMember) error {
	err := validateTeamMember("", teamMember)
	if err != nil {
		return err
	}
	return t.teamMembersRepo.UpdateTeamMember(teamMember)
}

func (t TeamMembersService) Delete(id string) error {
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return t.teamMembersRepo.DeleteTeamMember(uint32(idParsed))
}

func (t TeamMembersService) Find(id string) (domain.TeamMember, error) {
	idParsed, _ := strconv.Atoi(id)
	return t.teamMembersRepo.FindTeamMemberByID(uint32(idParsed))
}

func (t TeamMembersService) GetTeamMembersByMerchantCode(merchantCode string, page, pageSize int) ([]domain.TeamMember, error) {
	return t.teamMembersRepo.FindTeamMembersByMerchantCode(merchantCode, page, pageSize)
}

func validateTeamMember(action string, teamMember domain.TeamMember) error {
	switch strings.ToLower(action) {
	case "save":
		if teamMember.Name == "" {
			return errors.New("Required name")
		}
		if teamMember.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(teamMember.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	default:
		if teamMember.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(teamMember.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
