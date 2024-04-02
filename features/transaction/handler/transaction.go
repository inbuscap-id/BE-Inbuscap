package handler

import (
	"BE-Inbuscap/features/transaction"
	"BE-Inbuscap/helper"
	"net/http"
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
		if err := c.Bind(input); err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))

		}
		result, err := at.s.AddTransaction(c.Get("user").(*gojwt.Token), input.Amount)

		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorGeneralDatabase))

			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorDatabaseNotFound))

		}

		var response = new(TransactionRes)
		response.ID = result.ID
		response.OrderID = result.OrderID
		response.UserId = result.UserId
		response.Amount = result.Amount
		response.Status = result.Status
		response.Url = result.Url
		response.Token = result.Token
		response.CreatedAt = result.CreatedAt

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "transaction is created", response))
	}
}

// func (ct *TransactionHandler) CheckTransaction() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		transactionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 		if err != nil {
// 			return responses.PrintResponse(
// 				c, http.StatusBadRequest,
// 				"id tidak valid",
// 				nil)
// 		}
// 		result, err := ct.s.CheckTransaction(uint(transactionID))

// 		if err != nil {
// 			c.Logger().Error("Error fetching : ", err.Error())
// 			return responses.PrintResponse(
// 				c, http.StatusInternalServerError,
// 				"failed to retrieve data",
// 				nil)
// 		}

// 		var response = new(TransactionRes)
// 		response.ID = result.ID
// 		response.NoInvoice = result.NoInvoice
// 		response.JobID = result.JobID
// 		response.JobPrice = result.TotalPrice
// 		response.Status = result.Status

// 		return responses.PrintResponse(
// 			c, http.StatusOK,
// 			"transaction detail",
// 			response)
// 	}
// }

func (cb *TransactionHandler) CallBack() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CallBack)
		if err := c.Bind(input); err != nil {

			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))

		}
		result, err := cb.s.CallBack(input.OrderID)
		if err != nil {
			c.Logger().Error("something wrong: ", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorGeneralServer))

		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "midtrans callback successful", result))

	}
}
