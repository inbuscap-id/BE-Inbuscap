package handler

import (
	"BE-Inbuscap/features/transaction"
	"BE-Inbuscap/helper"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"

	echo "github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	s transaction.Service
}

func New(s transaction.Service) transaction.Controller {
	return &TransactionHandler{
		s: s,
	}
}

func (at *TransactionHandler) AddTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(TransactionReq)
		if err := c.Bind(&input); err != nil {
			log.Println(err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))

		}
		result, err := at.s.AddTransaction(c.Get("user").(*gojwt.Token), input.Amount)

		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorGeneralDatabase))

			}
			log.Println(err.Error())

			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorDatabaseNotFound))

		}

		var response = new(TransactionRes)
		response.ID = result.ID
		response.OrderID = result.OrderID
		response.UserId = result.UserId
		response.Amount = result.Amount
		response.Status = result.Status

		response.CreatedAt = result.CreatedAt

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "transaction is created", response))
	}
}
func (at *TransactionHandler) AddCoreTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(TransactionReq)
		if err := c.Bind(&input); err != nil {
			log.Println(err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))

		}
		result, err := at.s.AddCoreTransaction(c.Get("user").(*gojwt.Token), input.Amount, input.Bank)

		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralDatabase))

			}
			if strings.Contains(err.Error(), "midtrans") {
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "there's a problem in the server"))

			}
			log.Println(err.Error())

			return c.JSON(helper.ResponseFormat(http.StatusNotFound, helper.ErrorDatabaseNotFound))

		}
		// parsing result
		amt, _ := strconv.ParseFloat(result.GrossAmount, 64)
		var response = CoreTransactionRes{
			OrderID:    result.OrderID,
			Amount:     int(amt),
			Status:     result.TransactionStatus,
			CreatedAt:  result.TransactionTime,
			ExpiryTime: result.ExpiryTime,
		}

		if result.PermataVaNumber != "" {
			response.VaNumbers = map[string]string{
				"bank":      "permata",
				"va_number": result.PermataVaNumber,
			}
		} else {
			response.VaNumbers = result.VaNumbers
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "transaction is created", response))
	}
}

func (ct *TransactionHandler) CheckTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))

		}
		result, err := ct.s.CheckTransaction(uint(transactionID))

		if err != nil {
			c.Logger().Error("Error fetching : ", err.Error())
			log.Println(err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralDatabase))

		}

		var response = new(TransactionRes)
		response.ID = result.ID
		response.OrderID = result.OrderID
		response.UserId = result.UserId
		response.Amount = result.Amount
		response.Status = result.Status

		response.CreatedAt = result.CreatedAt

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "transaction is retireved", response))
	}
}

func (cb *TransactionHandler) CallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 1. Initialize empty map
		var notificationPayload map[string]interface{}

		// 2. Parse JSON request body and use it to set json to payload
		err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)
		if err != nil {
			// do something on error when decode
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}
		// 3. Get order-id from payload
		orderId, exists := notificationPayload["order_id"].(string)
		if !exists {
			// do something when key `order_id` not found
			log.Println("order id not found")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}
		result, err := cb.s.CallBack(orderId)
		if err != nil {
			c.Logger().Error("something wrong: ", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorGeneralServer))

		}
		var response = TransactionRes{
			ID:        result.ID,
			UserId:    result.UserId,
			OrderID:   result.OrderID,
			Amount:    result.Amount,
			Status:    result.Status,
			CreatedAt: result.CreatedAt,
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "midtrans callback successful", response))

	}
}
