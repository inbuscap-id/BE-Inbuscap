package midtrans

import (
	"BE-Inbuscap/config"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransCreateToken(orderID string, amount float64, namaCustomer string, email string, noHp string) *snap.Response {
	var s = snap.Client{}
	s.New(config.InitConfig().MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: namaCustomer,
			Email: email,
			Phone: noHp,
		},
		Items: &[]midtrans.ItemDetails{
			{
				Name:  "Top Up",
				Price: int64(amount),
				Qty:   1,
			},
		},
	}

	snapResp, _ := s.CreateTransaction(req)
	fmt.Println("snapresponse = ", snapResp)
	return snapResp
}

func MidtransStatus(orderID string) (Status string) {
	var c = coreapi.Client{}
	c.New(config.InitConfig().MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	transactionStatusResp, e := c.CheckTransaction(orderID)
	if e != nil {
		status := "Pending"
		return status
	} else {
		if transactionStatusResp != nil {
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					status := "Challenge"
					return status
				} else if transactionStatusResp.FraudStatus == "accept" {
					status := "Accept"
					return status
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				status := "Success"
				return status
			} else if transactionStatusResp.TransactionStatus == "deny" {
				status := "Deny"
				return status
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				status := "Cancelled"
				return status
			} else if transactionStatusResp.TransactionStatus == "pending" {
				status := "Pending"
				return status
			}
		}
	}

	status := "Pending"
	return status
}
