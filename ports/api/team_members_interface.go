package api

import (
	"merchant-account-api/core/domain"
)

type TeamMembersServiceInterface interface {
	Find(id string) (domain.TeamMember, error)
	Save(teamMember domain.TeamMember) error
	Update(teamMember domain.TeamMember) error
	Delete(id string) error
	GetTeamMembersByMerchantCode(merchantCode string, page, pageSize int) ([]domain.TeamMember, error)
}
