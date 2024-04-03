package data

import (
	"errors"

	"github.com/google/uuid"

	"BE-Inbuscap/features/transaction"

	"BE-Inbuscap/utils/midtrans"

	"gorm.io/gorm"
)

type TransactionQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.Repository {
	return &TransactionQuery{
		db: db,
	}
}

func (at *TransactionQuery) AddTransaction(userID uint, amount int) (transaction.Transaction, error) {
	var input Transaction
	input.UserID = userID
	input.Amount = amount
	input.Status = "pending"
	if err := at.db.Create(&input).Error; err != nil {
		return transaction.Transaction{}, err
	}

	newUUID := uuid.New()
	input.OrderID = string(newUUID.String())
	var usr = new(User)
	if err := at.db.First(usr, input.UserID).Error; err != nil {
		return transaction.Transaction{}, err
	}

	snap := midtrans.MidtransCreateToken(input.OrderID, amount, usr.Fullname, usr.Email, usr.Handphone)

	input.Url = snap.RedirectURL
	input.Token = snap.Token
	if err := at.db.Save(&input).Error; err != nil {
		return transaction.Transaction{}, err
	}

	var result transaction.Transaction
	result.ID = input.ID
	result.UserId = input.UserID
	result.Amount = input.Amount
	result.Status = input.Status
	result.Url = snap.RedirectURL
	result.Token = snap.Token
	result.OrderID = input.OrderID
	result.CreatedAt = input.CreatedAt

	return result, nil

}

func (ct *TransactionQuery) CheckTransaction(orderID string) (*transaction.Transaction, error) {
	var data Transaction
	if err := ct.db.Table("transactions").Where("order_id = ?", orderID).Find(&data).Error; err != nil {

		return &transaction.Transaction{}, err
	}

	if data.ID == 0 {
		err := errors.New("transaction doesnt exist")
		return nil, err
	}

	result := &transaction.Transaction{
		ID:        data.ID,
		UserId:    data.UserID,
		Amount:    data.Amount,
		Status:    data.Status,
		Url:       data.Url,
		Token:     data.Token,
		OrderID:   data.OrderID,
		CreatedAt: data.CreatedAt,
	}

	return result, nil

}

func (ct *TransactionQuery) CheckTransactionById(id uint) (*transaction.Transaction, error) {
	var data Transaction
	if err := ct.db.Table("transactions").Where(" id = ?", id).Find(&data).Error; err != nil {

		return &transaction.Transaction{}, err
	}

	if data.ID == 0 {
		err := errors.New("transaction doesnt exist")
		return nil, err
	}

	result := &transaction.Transaction{
		ID:        data.ID,
		UserId:    data.UserID,
		Amount:    data.Amount,
		Status:    data.Status,
		Url:       data.Url,
		Token:     data.Token,
		OrderID:   data.OrderID,
		CreatedAt: data.CreatedAt,
	}

	return result, nil
}

func (cb *TransactionQuery) Update(item transaction.Transaction) (*transaction.Transaction, error) {
	var data = Transaction{
		OrderID: item.OrderID,
		UserID:  item.UserId,
		Amount:  item.Amount,
		Status:  item.Status,
		Token:   item.Token,
		Url:     item.Url,
	}

	data.ID = item.ID

	var user = new(User)
	if err := cb.db.Where(" id = ?", data.UserID).First(&data).Error; err != nil {
		return nil, err
	}
	if data.Status == "Success" {
		user.Saldo += data.Amount
	}

	if err := cb.db.Save(&user).Error; err != nil {
		return nil, err
	}

	if err := cb.db.Save(&data).Error; err != nil {
		return nil, err
	}

	return &item, nil

}
