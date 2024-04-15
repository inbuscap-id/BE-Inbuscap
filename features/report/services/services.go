package services

import (
	"BE-Inbuscap/features/report"

	"github.com/go-playground/validator/v10"
)

type services struct {
	m report.Model
	v *validator.Validate
}

func Service(model report.Model) report.Services {
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
