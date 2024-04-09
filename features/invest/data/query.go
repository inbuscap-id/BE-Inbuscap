package data

import (
	"BE-Inbuscap/features/invest"
	"BE-Inbuscap/helper"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) invest.Model {
	return &model{
		connection: db,
	}
}

func (m *model) SendCapital(data invest.Investment) error {
	data.Status = 1
	err := m.connection.Create(&data).Error
	if err != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	}
	return nil
}

func (m *model) CancelSendCapital(data invest.Investment) error {
	var amount int
	m.connection.Table("investments").Select("SUM(amount) AS amount").Where("user_id = ? AND proposal_id = ?", data.User_id, data.Proposal_id).Scan(&amount)
	if amount <= 0 {
		return errors.New(helper.ErrorUserInput)
	}
	data.Amount = (-amount)
	data.Status = 2
	err := m.connection.Create(&data).Error
	if err != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	}
	return nil
}

func (m *model) GetAll(user_id uint, page int) (interface{}, int, error) {
	if page < 1 {
		page = 1
	}

	var result []struct {
		invest.Proposal
		Amount int
	}
	err := m.connection.Table("proposals").Select("proposals.*, SUM(investments.amount) AS collected, SUM(CASE WHEN investments.user_id = ? THEN investments.amount ELSE 0 END) AS amount", user_id).Group("proposals.id").Joins("JOIN investments ON investments.proposal_id = proposals.id").Having("amount <> 0").Limit(10).Offset(page*10 - 10).Scan(&result).Error

	var total_pages int
	m.connection.Table("investments").Select("COUNT(ID)").Group("proposal_id").Scan(&total_pages)
	return result, total_pages % 10, err
}

func (m *model) GetDetail(user_id uint, proposal_id int) (interface{}, error) {
	var result []invest.Investment
	err := m.connection.Where("user_id = ? AND proposal_id = ?", user_id, proposal_id).Find(&result).Error
	return result, err
}

// func (m *model) Edit(user_id string, proposal_id string, data invest.Investment) error {
// 	// if query := m.connection.Model(&data).Where("user_id = ? AND id = ?", user_id, proposal_id).Updates(&data); query.Error != nil {
// 	// 	return errors.New(helper.ErrorGeneralDatabase)
// 	// } else if query.RowsAffected == 0 {
// 	// 	return errors.New(helper.ErrorNoRowsAffected)
// 	// }
// 	return nil
// }

// func (m *model) Archive() error {
// 	return nil
// }
