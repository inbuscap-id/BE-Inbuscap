package services

import (
	"BE-Inbuscap/features/invest"
	"BE-Inbuscap/helper"
	"BE-Inbuscap/middlewares"
	"errors"
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

func (s *services) SendCapital(token *jwt.Token, proposal_id uint, amount int) error {
	// Take user_id from jwt
	id_string := middlewares.DecodeToken(token)

	// Convert user_id to uint
	user_id, err := strconv.ParseUint(id_string, 10, 32)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}

	// Check amount if less then 1
	if amount <= 0 {
		return errors.New(helper.ErrorUserInput)
	}

	// Collect data into one variable
	newInvest := invest.Investment{
		Proposal_id: proposal_id,
		User_id:     uint(user_id),
		Amount:      amount,
	}

	// Send data to database
	err = s.m.SendCapital(newInvest)
	if err != nil {
		return err
	}

	// Finish
	return nil
}

func (s *services) CancelSendCapital(token *jwt.Token, proposal_id string) error {
	// Take user_id from jwt
	id_string := middlewares.DecodeToken(token)

	// Convert user_id to uint
	user_id, err := strconv.ParseUint(id_string, 10, 32)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}

	// Convert page to uint
	proposalID, err := strconv.ParseUint(proposal_id, 10, 32)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}

	// Collect data into one variable
	newInvest := invest.Investment{
		Proposal_id: uint(proposalID),
		User_id:     uint(user_id),
	}

	// Send data to database
	err = s.m.CancelSendCapital(newInvest)

	// Finish
	return err
}

func (s *services) GetAll(token *jwt.Token, page string) (interface{}, int, error) {
	// Take user_id from jwt
	id_string := middlewares.DecodeToken(token)

	// Convert user_id to uint
	user_id, err := strconv.ParseUint(id_string, 10, 32)
	if err != nil {
		return nil, 0, errors.New(helper.ErrorUserInput)
	}

	// Convert page to int
	page_int := 1
	if page != "" {
		page_int, err = strconv.Atoi(page)
		if err != nil {
			return nil, 0, errors.New(helper.ErrorUserInput)
		}
	}

	// Get data from database
	data, total_pages, err := s.m.GetAll(uint(user_id), page_int)
	if err != nil {
		return nil, 0, errors.New(helper.ErrorGeneralDatabase)
	}

	if reflect.ValueOf(data).Len() == 0 {
		return []struct{}{}, 0, nil
	}

	// Finish
	return data, total_pages, nil
}

func (s *services) GetDetail(token *jwt.Token, proposal_id string) (interface{}, error) {
	// Take user_id from jwt
	id_string := middlewares.DecodeToken(token)

	// Convert user_id to uint
	user_id, err := strconv.ParseUint(id_string, 10, 32)
	if err != nil {
		return nil, errors.New(helper.ErrorUserInput)
	}

	// Convert proposal_id to int
	proposalID, err := strconv.Atoi(proposal_id)
	if err != nil {
		return nil, errors.New(helper.ErrorUserInput)
	}

	// Get data from database
	data, err := s.m.GetDetail(uint(user_id), proposalID)
	if err != nil {
		return nil, errors.New(helper.ErrorGeneralDatabase)
	}

	// Finish
	return data, nil
}

// func (s *services) Edit() error {
// 	return nil
// }

// func (s *services) Archive() error {
// 	return nil
// }
