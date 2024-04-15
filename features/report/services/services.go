package services

import (
	"BE-Inbuscap/features/report"
	"BE-Inbuscap/helper"
	utils "BE-Inbuscap/utils/cloudinary"
	"errors"
	"mime/multipart"
	"strconv"

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

func (s *services) AddReport(token interface{}, proposal_id string, document *multipart.FileHeader) error {
	proposalID, err := strconv.ParseUint(proposal_id, 10, 32)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}

	documentURL, err := utils.UploadImage(document)
	if err != nil {
		return err
	}

	data_report := report.Report{
		Proposal_id: uint(proposalID),
		Document:    documentURL,
	}

	return s.m.AddReport(data_report)
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
