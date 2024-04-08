package handler

import (
	"BE-Inbuscap/features/invest"
	"BE-Inbuscap/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s invest.Services
}

func NewHandler(service invest.Services) invest.Controller {
	return &controller{
		s: service,
	}
}

func (ct *controller) SendCapital() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input InvestmentRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err = ct.s.SendCapital(token, input.Proposal_id, input.Amount)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Sent Capital"))
	}
}

func (ct *controller) CancelSendCapital() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err := ct.s.CancelSendCapital(token, c.QueryParam("proposal_id"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Cancel The Sent Capital"))
	}
}

func (ct *controller) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		data, total_pages, err := ct.s.GetAll(token, c.QueryParam("page"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		var dataResponse []InvestmentResponse
		helper.ConvertStruct(&data, &dataResponse)

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Success Get All Investments", dataResponse,
			map[string]any{
				"pagination": map[string]any{
					"page": func(p string) int {
						page, _ := strconv.Atoi(p)
						if page <= 0 {
							page = 1
						}
						return page
					}(c.QueryParam("page")),
					"page_size":   10,
					"total_pages": total_pages,
				},
			},
		))
	}
}

func (ct *controller) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		data, err := ct.s.GetDetail(token, c.Param("proposal_id"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		var dataResponse []InvestmentDetileResponse
		helper.ConvertStruct(&data, &dataResponse)

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Get Detail Investment Proposal", dataResponse))
	}
}

// func (ct *controller) Edit() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(helper.ResponseFormat(http.StatusCreated, "Success create post", nil))
// 	}
// }

// func (ct *controller) Archive() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(helper.ResponseFormat(http.StatusCreated, "Success create post", nil))
// 	}
// }
