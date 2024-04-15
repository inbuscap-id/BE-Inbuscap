package data

import (
	entity "BE-Inbuscap/features/report"

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

func (m *model) AddReport(data entity.Report) error {
	return m.connection.Create(&data).Error
}

func (m *model) Edit() error {
	return nil
}

func (m *model) GetAll() error {
	return nil
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
