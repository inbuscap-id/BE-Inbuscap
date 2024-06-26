package data

import (
	"BE-Inbuscap/features/proposal"
	"BE-Inbuscap/helper"
	"errors"
	"log"

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
	return m.connection.Create(&data).Error
}

func (m *model) Update(user_id string, proposal_id string, data proposal.Proposal) error {
	var selectUpdate []string
	if data.Title != "" {
		selectUpdate = append(selectUpdate, "title")
	}
	if data.Image != "" {
		selectUpdate = append(selectUpdate, "image")
	}
	if data.Document != "" {
		selectUpdate = append(selectUpdate, "document")
	}
	if data.Description != "" {
		selectUpdate = append(selectUpdate, "description")
	}
	if data.Capital != 0 {
		selectUpdate = append(selectUpdate, "capital")
	}
	if data.Share != 0 {
		selectUpdate = append(selectUpdate, "share")
	}
	if len(selectUpdate) == 0 {
		return errors.New(helper.ErrorNoRowsAffected)
	}
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
	err := m.connection.Select("proposals.*, SUM(investments.amount) AS collected").Group("proposals.id").Joins("LEFT JOIN investments ON investments.proposal_id = proposals.id").Order("proposals.created_at DESC").Limit(10).Offset(page*10 - 10).Find(&result).Error

	var numberOfProposals int
	m.connection.Table("proposals").Select("COUNT(ID)").Where("`proposals`.`deleted_at` IS NULL").Scan(&numberOfProposals)
	return result, (numberOfProposals + 9) / 10, err
}

func (m *model) GetAllMy(page int, user_id string) ([]proposal.Proposal, int, error) {
	if page < 1 {
		page = 1
	}
	var result []proposal.Proposal
	err := m.connection.Select("proposals.*, SUM(investments.amount) AS collected").Group("proposals.id").Joins("LEFT JOIN investments ON investments.proposal_id = proposals.id").Where("proposals.user_id = ?", user_id).Order("proposals.created_at DESC").Limit(10).Offset(page*10 - 10).Find(&result).Error

	var numberOfProposals int
	m.connection.Table("proposals").Select("COUNT(ID)").Where("`proposals`.`deleted_at` IS NULL").Scan(&numberOfProposals)
	return result, (numberOfProposals + 9) / 10, err
}

func (m *model) GetDetail(id_proposal string) (proposal.Proposal, error) {
	var result proposal.Proposal
	err := m.connection.Select("proposals.*, SUM(investments.amount) AS collected").Preload("User").Joins("LEFT JOIN investments ON investments.proposal_id = proposals.id").Where("proposals.id = ?", id_proposal).Find(&result).Error
	return result, err
}

func (m *model) Delete(id string, prososal_id string) error {
	if query := m.connection.Where("user_id = ? AND id = ?", id, prososal_id).Delete(&proposal.Proposal{}); query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorDatabaseNotFound)
	}
	return nil
}

func (m *model) Archive() error {
	return nil
}

func (m *model) GetVerifications(page int, status int) ([]proposal.Proposal, int, []string, error) {
	if page < 1 {
		page = 1
	}
	if status < 0 {
		status = 0
	}
	var result []proposal.Proposal
	var total []proposal.Proposal
	err := m.connection.Order("updated_at desc").Where(" status = ? AND deleted_at IS NULL", status).
		Find(&total).Error
	if err != nil {
		log.Println("error mengambil proposal", err.Error())
		return nil, 0, nil, err

	}
	err = m.connection.Order("updated_at desc").Where(" status = ? AND deleted_at IS NULL", status).
		Limit(10).Offset(page*10 - 10).Find(&result).Error

	if err != nil {
		log.Println("error mengambil proposal", err.Error())
		return nil, 0, nil, err

	}
	var users []string
	for _, val := range result {
		var user proposal.User
		err = m.connection.Where("id = ?", val.User_id).First(&user).Error
		if err != nil {
			log.Println("error mengambil user", err.Error())
			return nil, 0, nil, err
		}
		users = append(users, user.Fullname)
	}
	return result, (len(total) + 9) / 10, users, nil
}

func (m *model) ChangeStatus(id uint, status int) error {
	var result proposal.Proposal
	err := m.connection.Where("id = ?", id).First(&result).Error
	if err != nil {
		log.Println("error mengambil data", err.Error())
		return err
	}
	result.Status = status
	err = m.connection.Save(&result).Error

	return err
}

func (m *model) GetVerification(proposalId uint) (proposal.Proposal, string, error) {
	var result proposal.Proposal
	err := m.connection.Where("id = ?", proposalId).First(&result).Error
	if err != nil {
		log.Println("error mengambil proposal", err.Error())
		return proposal.Proposal{}, "", err
	}
	var user proposal.User
	err = m.connection.Where("id = ?", result.User_id).First(&user).Error
	if err != nil {
		log.Println("error mengambil user", err.Error())
		return proposal.Proposal{}, "", err
	}
	return result, user.Fullname, nil
}
