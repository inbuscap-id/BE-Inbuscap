package services

import (
	"BE-Inbuscap/features/invest"

	"github.com/go-playground/validator/v10"
)

type services struct {
	m invest.Model
	v *validator.Validate
}

func Service(model invest.Model) invest.Services {
	return &services{
		m: model,
		v: validator.New(),
	}
}

func (s *services) Create() error {
	return nil
}

func (s *services) Edit() error {
	return nil
}

func (s *services) GetAll() error {
	return nil
}

func (s *services) GetDetail() error {
	return nil
}

func (s *services) Delete() error {
	return nil
}

func (s *services) Archive() error {
	return nil
}
