package data

import (
	"BE-Inbuscap/features/proposal"
	"BE-Inbuscap/helper"
	"errors"
	"time"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) proposal.Model {
	return &model{
		connection: db,
	}
}

func (m *model) Create(data proposal.Proposal) error {
	data.CreatedAt = time.Now().UTC()
	data.UpdatedAt = time.Now().UTC()
	return m.connection.Create(&data).Error
}

func (m *model) Update(user_id string, proposal_id string, data proposal.Proposal) error {
	data.UpdatedAt = time.Now().UTC()
	if query := m.connection.Model(&data).Where("user_id = ? AND id = ?", user_id, proposal_id).Updates(&data); query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorNoRowsAffected)
	}
	return nil
}

func (m *model) GetAll(page int) ([]proposal.Proposal, int, error) {
	if page < 1 {
		page = 1
	}
	var result []proposal.Proposal
	err := m.connection.Limit(10).Offset(page*10 - 10).Find(&result).Error

	var numberOfProposals int
	m.connection.Table("Proposals").Select("COUNT(ID)").Where("proposal_id = ?", 1).Scan(&numberOfProposals)
	return result, numberOfProposals % 10, err
}

func (m *model) GetDetail(id_proposal string) (proposal.Proposal, error) {
	var result proposal.Proposal
	err := m.connection.Joins("User").Where("ID = ?", id_proposal).Find(&result).Error
	return result, err
}

func (m *model) Delete(id string, prososal_id string) error {
	if query := m.connection.Where("user_id AND id = ?", id, prososal_id).Delete(&proposal.Proposal{}); query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorDatabaseNotFound)
	}
	return nil
}

func (m *model) Archive() error {
	return nil
}
