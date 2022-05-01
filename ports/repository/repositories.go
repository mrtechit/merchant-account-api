package repository

import (
	"merchant-account-api/core/domain"
)

type TeamMembersRepository interface {
	FindTeamMemberByID(id uint32) (domain.TeamMember, error)
	FindTeamMembersByMerchantCode(merchantCode string, page, pageSize int) ([]domain.TeamMember, error)
	SaveTeamMember(teamMember domain.TeamMember) error
	UpdateTeamMember(teamMember domain.TeamMember) error
	DeleteTeamMember(id uint32) error
}

type MerchantRepository interface {
	FindMerchantByID(id string) (domain.Merchant, error)
	SaveMerchant(merchant domain.Merchant) (string, error)
	UpdateMerchant(merchant domain.Merchant) error
	DeleteMerchant(id string) error
}
