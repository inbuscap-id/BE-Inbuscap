package services

import (
	"BE-Inbuscap/features/proposal"
	"BE-Inbuscap/helper"
	"BE-Inbuscap/middlewares"
	utils "BE-Inbuscap/utils/cloudinary"
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type services struct {
	m proposal.Model
	v *validator.Validate
}

func Service(model proposal.Model) proposal.Services {
	return &services{
		m: model,
		v: validator.New(),
	}
}

func (s *services) Create(token *jwt.Token, image *multipart.FileHeader, document *multipart.FileHeader, data proposal.Proposal) error {
	id_string := middlewares.DecodeToken(token)

	id, _ := strconv.ParseUint(id_string, 10, 32)

	imageURL, err := utils.UploadImage(image)
	if err != nil {
		return err
	}

	documentURL, err := utils.UploadImage(document)
	if err != nil {
		return err
	}

	data.User_id = uint(id)
	data.Image = imageURL
	data.Document = documentURL

	return s.m.Create(data)
}

func (s *services) Update(token *jwt.Token, proposal_id string, image *multipart.FileHeader, document *multipart.FileHeader, data proposal.Proposal) error {
	user_id := middlewares.DecodeToken(token)

	_, err := strconv.Atoi(proposal_id)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}

	imageURL, err := utils.UploadImage(image)
	if err != nil {
		return err
	}

	documentURL, err := utils.UploadImage(document)
	if err != nil {
		return err
	}

	data.Image = imageURL
	data.Document = documentURL

	return s.m.Update(user_id, proposal_id, data)
}

func (s *services) GetAll(page string) ([]proposal.Proposal, int, error) {
	var page_int = 0
	page_int, err := strconv.Atoi(page)
	if page != "" && err != nil {
		return []proposal.Proposal{}, 0, errors.New(helper.ErrorUserInput)
	}

	listProposal, totalPage, err := s.m.GetAll(page_int)
	if err != nil {
		return []proposal.Proposal{}, 0, errors.New(helper.ErrorGeneralDatabase)
	}

	if len(listProposal) == 0 {
		return []proposal.Proposal{}, 0, nil
	}

	return listProposal, totalPage, nil
}

func (s *services) GetDetail(id_proposal string) (proposal.Proposal, error) {
	detileProposal, err := s.m.GetDetail(id_proposal)

	if detileProposal.Title == "" && detileProposal.Capital == 0 {
		return proposal.Proposal{}, errors.New(helper.ErrorDatabaseNotFound)
	}

	if err != nil {
		return proposal.Proposal{}, errors.New(helper.ErrorGeneralDatabase)
	}

	return detileProposal, nil
}

func (s *services) Delete(token *jwt.Token, prososal_id string) error {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Delete Date
	if err := s.m.Delete(decodeID, prososal_id); err != nil {
		return err
	}

	// Finished
	return nil
}

func (s *services) Archive() error {
	return nil
}
