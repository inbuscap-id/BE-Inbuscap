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
			log.Println("error file data:", err.Error())
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
		share, err := strconv.Atoi(c.FormValue("share"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		newProposal := proposal.Proposal{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Capital:     capital,
			Share:       share,
		}

		err = ct.s.Create(token, fileImage, fileProposal, newProposal)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success create post", nil))
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
		return c.JSON(helper.ResponseFormat(http.StatusOK, "success update post", nil))
	}
}

func (ct *controller) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, total_pages, err := ct.s.GetAll(c.QueryParam("page"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		var dataResponse []ProposalResponse
		helper.ConvertStruct(&data, &dataResponse)

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Get All Proposals", dataResponse,
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

func (ct *controller) GetAllMy() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		data, total_pages, err := ct.s.GetAllMy(token, c.QueryParam("page"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		var dataResponse []ProposalResponse
		helper.ConvertStruct(&data, &dataResponse)

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Get All MyProposals", dataResponse,
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
		data, err := ct.s.GetDetail(c.Param("proposal_id"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		var dataResponse ProposalDetailResponse
		helper.ConvertStruct(&data, &dataResponse)

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success get detail post", dataResponse))
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

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success delete post", nil))
	}
}

func (ct *controller) Archive() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(helper.ResponseFormat(http.StatusOK, "success archive post", nil))
	}
}

func (ct *controller) GetVerifications() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}
		status, err := strconv.Atoi(c.QueryParam("status"))
		if err != nil || status <= 0 {
			status = 0
		}
		data, total_pages, users, err := ct.s.GetVerifications(page, status)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}
		var paginasi helper.Pagination
		paginasi.Page = page
		paginasi.PageSize = 10

		var dataResponse []VerificationResponse
		helper.ConvertStruct(&data, &dataResponse)

		for i := range dataResponse {
			dataResponse[i].Owner = users[i]
			dataResponse[i].Document = data[i].Document
		}

		paginasi.TotalPages = total_pages
		if page > total_pages {
			return c.JSON(helper.ResponseFormat(http.StatusNotFound, "index out of bounds"))

		}
		return c.JSON(helper.ResponseFormatArray(http.StatusOK, "data retrieved successfully", dataResponse, paginasi))
	}
}

func (ct *controller) ChangeStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		stringID, err := strconv.Atoi(c.Param("proposal_id"))
		if err != nil {
			log.Println("error mengambil user id,", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))

		}
		id := uint(stringID)
		var Req ChangeStatus
		err = c.Bind(&Req)
		if err != nil {
			log.Println("error mengambil req body,", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))

		}
		err = ct.s.ChangeStatus(id, int(Req.Status))
		if err != nil {
			log.Println("error saat mengubah data,", err.Error())
			if strings.Contains(err.Error(), "found") {
				return c.JSON(helper.ResponseFormat(http.StatusNotFound, helper.ErrorDatabaseNotFound))

			}
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralServer))

		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Updated"))

	}
}
