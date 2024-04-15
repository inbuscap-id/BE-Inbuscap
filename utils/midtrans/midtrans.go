package midtrans

import (
	"BE-Inbuscap/config"
	"log"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransTokenCore(bank string, amount int) (*coreapi.ChargeResponse, *midtrans.Error) {
	cfg := config.InitConfig()
	midtrans.ServerKey = cfg.MIDTRANS_SERVER_KEY
	midtrans.Environment = midtrans.Sandbox
	log.Println(cfg.MIDTRANS_SERVER_KEY)
	newUUID := uuid.New()
	order := string(newUUID.String())

	chargReq := &coreapi.ChargeReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order,
			GrossAmt: int64(amount),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(bank),
		},
		PaymentType: coreapi.CoreapiPaymentType("bank_transfer"),
		Items: &[]midtrans.ItemDetails{
			{
				Name:  "Top Up",
				Price: int64(amount),
				Qty:   1,
			},
		},
	}

	resp, err := coreapi.ChargeTransaction(chargReq)
	if err != nil {
		log.Println("midtrans error", err.Error())
		return nil, err
	}
	return resp, nil
}
func MidtransCreateToken(orderID string, amount int, namaCustomer string, email string, noHp string) *snap.Response {
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
	return snapResp
}

func MidtransStatus(orderID string) (Status string) {
	cfg := config.InitConfig()
	midtrans.ServerKey = cfg.MIDTRANS_SERVER_KEY
	midtrans.Environment = midtrans.Sandbox
	log.Println(cfg.MIDTRANS_SERVER_KEY)

	transactionStatusResp, e := coreapi.CheckTransaction(orderID)
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
