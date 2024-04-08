package handler

import (
	"BE-Inbuscap/features/user"
	"BE-Inbuscap/helper"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service user.Service
}

func NewHandler(s user.Service) user.Controller {
	return &controller{
		service: s,
	}
}

func (ct *controller) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		var data RegisterRequest
		err := c.Bind(&data)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}
		var input = user.Register{
			Fullname:  data.Fullname,
			Email:     data.Email,
			Handphone: data.Handphone,
			KTP:       data.KTP,
			NPWP:      data.NPWP,
			Password:  data.Password,
		}
		err = ct.service.Register(input)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "Registered Successfully"))
	}
}

func (ct *controller) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		var processData user.User
		processData.Email = input.Email
		processData.Password = input.Password

		usertoken, err := ct.service.Login(processData)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Login Successfully", map[string]any{"token": usertoken}))
	}
}

func (ct *controller) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		profile, err := ct.service.Profile(token)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralServer))
		}

		var profileResponse = ProfileResponse{

			Fullname:  profile.Fullname,
			Email:     profile.Email,
			Handphone: profile.Handphone,
			KTP:       profile.KTP,
			NPWP:      profile.NPWP,
			PhotoKTP:  profile.PhotoKTP,
			PhotoNPWP: profile.PhotoNPWP,
			PhotoSelf: profile.PhotoSelf,
			Avatar:    profile.Avatar,
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Get MyProfile", profileResponse))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input user.User
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

		err = ct.service.Update(token, input)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Updated"))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err := ct.service.Delete(token)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Deleted User"))
	}
}

func (ct *controller) AddVerification() echo.HandlerFunc {
	return func(c echo.Context) error {
		ktp, err := helper.SelectFile(c, "photo_ktp")
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInputFormat))
		}
		npwp, err := helper.SelectFile(c, "photo_npwp")
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInputFormat))
		}
		selfie, err := helper.SelectFile(c, "photo_selfie")
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInputFormat))
		}
		var input = []*multipart.FileHeader{}
		input = append(input, ktp, npwp, selfie)

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			log.Println("error saat mengambil token")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}
		// log.Println("")

		err = ct.service.AddVerification(token, input)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "Successfully Updated"))
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

		var paginasi helper.Pagination
		paginasi.Page = page
		paginasi.PageSize = 10
		result, count, err := ct.service.GetVerifications(paginasi, status)
		if err != nil {
			log.Println("error handler", err.Error())
		}
		totalPages := int(math.Ceil(float64(count) / float64(paginasi.PageSize)))
		paginasi.TotalPages = totalPages
		var payloads []GetVerificationsResponse
		for _, val := range result {
			var payload = GetVerificationsResponse{
				Fullname:  val.Fullname,
				Handphone: val.Handphone,
				KTP:       val.KTP,
				NPWP:      val.NPWP,
				PhotoKTP:  val.PhotoKTP,
				PhotoNPWP: val.PhotoNPWP,
				PhotoSelf: val.PhotoSelf,
				IsActive:  val.IsActive,
			}
			payload.ID = val.ID
			payloads = append(payloads, payload)
		}

		return c.JSON(helper.ResponseFormatArray(http.StatusOK, "User list sucessfully retrieved", payloads, paginasi))
	}
}
