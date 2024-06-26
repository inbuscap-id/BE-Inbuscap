package services

import (
	"BE-Inbuscap/features/transaction"
	jwt "BE-Inbuscap/middlewares"
	"BE-Inbuscap/utils/midtrans"
	"errors"
	"log"
	"strconv"

	golangjwt "github.com/golang-jwt/jwt/v5"
	"github.com/midtrans/midtrans-go/coreapi"
)

type TransactionService struct {
	repo transaction.Repository
}

func New(r transaction.Repository) transaction.Service {
	return &TransactionService{
		repo: r,
	}
}

func (at *TransactionService) AddTransaction(token *golangjwt.Token, amount int) (transaction.Transaction, error) {
	userID := jwt.DecodeToken(token)
	id, err := strconv.Atoi(userID)
	if err != nil {
		return transaction.Transaction{}, err
	}

	result, err := at.repo.AddTransaction(uint(id), amount)
	if err != nil {
		log.Println(err.Error())
	}
	return result, err
}

func (at *TransactionService) AddCoreTransaction(token *golangjwt.Token, amount int, bank string) (*coreapi.ChargeResponse, error) {
	userID := jwt.DecodeToken(token)
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("error convert id", err.Error())
		return nil, err
	}
	resp, er := midtrans.MidtransTokenCore(bank, amount)
	if er != nil {
		log.Println(er.Error())
		return nil, errors.New("error di midtrans")
	}
	result, err := at.repo.AddCoreTransaction(uint(id), resp)
	if err != nil || result.ID == 0 {
		log.Println(err.Error())
		return nil, err
	}
	return resp, err
}

func (ct *TransactionService) CheckTransaction(transactionID uint) (transaction.Transaction, error) {
	result, err := ct.repo.CheckTransactionById(transactionID)
	if err != nil {
		return transaction.Transaction{}, err
	}
	return *result, err
}

func (cb *TransactionService) CallBack(orderID string) (transaction.Transaction, error) {
	data, err := cb.repo.CheckTransaction(orderID)
	if err != nil {
		log.Println(err.Error())
		return transaction.Transaction{}, err
	}
	ms := midtrans.MidtransStatus(orderID)
	data.Status = ms
	result, err := cb.repo.Update(*data)
	if err != nil {
		log.Println(err.Error())

		return transaction.Transaction{}, err
	}
	if result == nil {
		return transaction.Transaction{}, errors.New("result is nil")
	}
	return *result, err
}
