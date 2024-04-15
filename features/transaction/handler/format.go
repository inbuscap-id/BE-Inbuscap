package handler

import "time"

type TransactionReq struct {
	Amount int    `json:"amount"`
	Bank   string `json:"bank"`
}

type TransactionRes struct {
	ID      uint   `json:"transaction_id"`
	OrderID string `json:"order_id"`
	UserId  uint   `json:"user_id"`
	Amount  int    `json:"amount"`
	Status  string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}
type CoreTransactionRes struct {
	OrderID   string      `json:"order_id"`
	Amount    int         `json:"amount"`
	Status    string      `json:"status"`
	VaNumbers interface{} `json:"va_numbers"`
	CreatedAt string      `json:"created_at"`
}

type CallBack struct {
	OrderID string `json:"order_id"`
}
