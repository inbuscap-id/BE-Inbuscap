package handler

import (
	"BE-Inbuscap/features/user"
	"BE-Inbuscap/helper"
	"net/http"
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

		var profileResponse ProfileResponse
		helper.ConvertStruct(&profile, &profileResponse)

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
