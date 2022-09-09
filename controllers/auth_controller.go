package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/saintox/go-basic-auth/constants"
	"github.com/saintox/go-basic-auth/dtos"
	"github.com/saintox/go-basic-auth/entities"
	"github.com/saintox/go-basic-auth/middlewares"
	"github.com/saintox/go-basic-auth/pkg/logger"
	"github.com/saintox/go-basic-auth/services"
	"net/http"
)

type AuthController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

type AuthControllerImpl struct {
	validator *middlewares.CustomValidator
	service   *services.Service
	logger    *logger.Logger
}

func NewAuthController(service *services.Service, validator *middlewares.CustomValidator, logger *logger.Logger) *AuthControllerImpl {
	return &AuthControllerImpl{
		service:   service,
		validator: validator,
		logger:    logger,
	}
}

func (ctl AuthControllerImpl) Login(ctx echo.Context) (err error) {
	var (
		requestParam dtos.LoginRequestBody
	)

	if err = ctx.Bind(&requestParam); err != nil {
		data := dtos.NewResponse(http.StatusBadRequest, constants.FailedParseRequest, err, ctl.logger)
		return ctx.JSON(http.StatusBadRequest, data)
	}

	if err = ctl.validator.Validate(&requestParam); err != nil {
		data := dtos.NewResponse(http.StatusUnprocessableEntity, constants.ValidationError, err, ctl.logger)
		return ctx.JSON(http.StatusUnprocessableEntity, data)
	}

	check, err := ctl.service.Login.CheckCredential(requestParam.Email, requestParam.Password)
	if err != nil {
		data := dtos.NewResponse(http.StatusUnauthorized, constants.InvalidCredential, err, ctl.logger)
		return ctx.JSON(http.StatusUnauthorized, data)
	}

	return ctx.JSON(http.StatusOK, check)
}

func (ctl AuthControllerImpl) Register(ctx echo.Context) (err error) {
	var requestParam dtos.RegisterRequestBody

	if err = ctx.Bind(&requestParam); err != nil {
		data := dtos.NewResponse(http.StatusBadRequest, constants.FailedParseRequest, err, ctl.logger)
		return ctx.JSON(http.StatusBadRequest, data)
	}

	if err = ctl.validator.Validate(&requestParam); err != nil {
		data := dtos.NewResponse(http.StatusUnprocessableEntity, constants.ValidationError, err, ctl.logger)
		return ctx.JSON(http.StatusUnprocessableEntity, data)
	}

	createUser := &entities.User{
		Name:     requestParam.Name,
		Email:    requestParam.Email,
		Password: requestParam.Password,
	}

	create, err := ctl.service.Register.UserRegister(createUser)
	if err != nil {
		data := dtos.NewResponse(http.StatusBadRequest, constants.FailedParseRequest, err, ctl.logger)
		return ctx.JSON(http.StatusBadRequest, data)
	}

	return ctx.JSON(http.StatusCreated, create)
}
