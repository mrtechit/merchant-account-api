package sql_repository

import (
	"gorm.io/gorm"
	"merchant-account-api/core/domain"
)

func (p *postgreSQLRepository) FindTeamMemberByID(id uint32) (domain.TeamMember, error) {

	teamMember := domain.TeamMember{}
	result := p.db.First(&teamMember, "id = ?", id)
	if result.Error != nil {
		return teamMember, result.Error
	}

	return teamMember, nil
}

func (p *postgreSQLRepository) SaveTeamMember(teamMember domain.TeamMember) error {
	teamMemberSave := domain.TeamMember{
		Name:                 teamMember.Name,
		Email:                teamMember.Email,
		MerchantMerchantCode: teamMember.MerchantMerchantCode,
	}
	result := p.db.Create(&teamMemberSave)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *postgreSQLRepository) UpdateTeamMember(teamMember domain.TeamMember) error {
	result := p.db.Model(&teamMember).Where("id = ?", teamMember.ID).Updates(domain.TeamMember{
		Name:  teamMember.Name,
		Email: teamMember.Email,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *postgreSQLRepository) DeleteTeamMember(id uint32) error {
	result := p.db.Where("id = ?", id).Delete(&domain.TeamMember{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *postgreSQLRepository) FindTeamMembersByMerchantCode(merchantCode string, page, pageSize int) ([]domain.TeamMember, error) {
	var teamMembers []domain.TeamMember
	merchant := domain.Merchant{}
	result := p.db.First(&merchant, "merchant_code = ?", merchantCode)
	if result.Error != nil {
		return teamMembers, result.Error
	}

	p.db.Preload("Merchant", "merchant_code = ?", merchantCode).Scopes(Paginate(page, pageSize)).Find(&teamMembers)
	return teamMembers, nil
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
