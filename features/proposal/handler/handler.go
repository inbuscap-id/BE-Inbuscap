package handler

import (
	"BE-Inbuscap/features/proposal"
	"BE-Inbuscap/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s proposal.Services
}

func NewHandler(service proposal.Services) proposal.Controller {
	return &controller{
		s: service,
	}
}

func (ct *controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileImage, err := c.FormFile("image")
		if err != nil {
			log.Println("error image data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat, nil))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput, nil))
		}

		fileProposal, err := c.FormFile("proposal")
		if err != nil {
			log.Println("error image data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat, nil))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		capital, err := strconv.Atoi(c.FormValue("capital"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		newProposal := proposal.Proposal{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Capital:     capital,
		}

		err = ct.s.Create(token, fileImage, fileProposal, newProposal)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileImage, err := c.FormFile("image")
		if err != nil {
			log.Println("error image data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat, nil))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput, nil))
		}

		fileProposal, err := c.FormFile("proposal")
		if err != nil {
			log.Println("error image data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat, nil))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		capital, err := strconv.Atoi(c.FormValue("capital"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		dataProposal := proposal.Proposal{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Capital:     capital,
		}

		err = ct.s.Update(token, c.Param("proposal_id"), fileImage, fileProposal, dataProposal)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success update post", nil))
	}
}

func (ct *controller) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		proposals, totalPage, err := ct.s.GetAll(c.QueryParam("page"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success get posts", proposals, map[string]any{"total_pages": totalPage}))
	}
}

func (ct *controller) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		detailProposal, err := ct.s.GetDetail(c.Param("proposal_id"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success get detail post", detailProposal))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err := ct.s.Delete(token, c.Param("proposal_id"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success delete post", nil))
	}
}

func (ct *controller) Archive() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success archive post", nil))
	}
}
