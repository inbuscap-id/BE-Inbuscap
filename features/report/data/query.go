package data

import (
	entity "BE-Inbuscap/features/report"
	"BE-Inbuscap/helper"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) entity.Model {
	return &model{
		connection: db,
	}
}

func (m *model) AddReport(data entity.Report, userID string) error {
	var userID_proposal string
	m.connection.Table("proposals").Select("proposals.user_id as userID_proposal").
		Where("proposals.id = ?", data.Proposal_id).Scan(&userID_proposal)
	if userID_proposal != userID {
		return errors.New(helper.ErrorUserInput)
	}
	return m.connection.Create(&data).Error
}

func (m *model) Edit() error {
	return nil
}

func (m *model) GetAllReport(proposal_id string) (entity.Proposal, error) {
	var proposal entity.Proposal
	err := m.connection.Preload("reports").Where("id = ?", proposal_id).Find(&proposal).Error
	if err != nil {
		return entity.Proposal{}, err
	}
	return proposal, nil
}

func (m *model) GetDetail() error {
	return nil
}

func (m *model) Delete() error {
	return nil
}

func (m *model) Archive() error {
	return nil
}
