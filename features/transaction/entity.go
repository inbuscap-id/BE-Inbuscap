package transaction

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Transaction struct {
	ID        uint
	OrderID   string
	UserId    uint
	Amount    int
	Status    string
	Token     string
	Url       string
	CreatedAt time.Time `json:"created_at"`
}

// type TransactionList struct {
// 	ID         uint
// 	NoInvoice  string
// 	JobID      uint
// 	TotalPrice uint
// 	Status     string
// 	Token      string
// 	Url        string
// 	Timestamp  time.Time `json:"timestamp"`
// }

type Controller interface {
	AddTransaction() echo.HandlerFunc
	AddCoreTransaction() echo.HandlerFunc
	CheckTransaction() echo.HandlerFunc
	CallBack() echo.HandlerFunc
}

type Repository interface {
	AddTransaction(userID uint, amount int) (Transaction, error)
	AddCoreTransaction(userID uint, transaksi *coreapi.ChargeResponse) (Transaction, error)

	CheckTransaction(orderID string) (*Transaction, error)
	CheckTransactionById(id uint) (*Transaction, error)
	Update(item Transaction) (*Transaction, error)
}

type Service interface {
	AddTransaction(token *jwt.Token, amount int) (Transaction, error)
	AddCoreTransaction(token *jwt.Token, amount int, bank string) (*coreapi.ChargeResponse, error)

	CheckTransaction(transactionID uint) (Transaction, error)
	CallBack(noInvoice string) (Transaction, error)
}
