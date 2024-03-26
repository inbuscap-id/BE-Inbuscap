package data

import (
	"BE-Inbuscap/features/invest"

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

func (pm *model) Create() error {
	return nil
}

func (pm *model) Edit() error {
	return nil
}

func (pm *model) GetAll() error {
	return nil
}

func (pm *model) GetDetail() error {
	return nil
}

func (pm *model) Delete() error {
	return nil
}

func (pm *model) Archive() error {
	return nil
}
