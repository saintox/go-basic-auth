package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/saintox/go-basic-auth/dtos"
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
	validator *middlewares.Validator
	service   *services.Service
	logger    *logger.Logger
}

func NewAuthController(service *services.Service, validator *middlewares.Validator, logger *logger.Logger) *AuthControllerImpl {
	return &AuthControllerImpl{
		service:   service,
		validator: validator,
		logger:    logger,
	}
}

func (ctl AuthControllerImpl) Login(ctx echo.Context) error {
	var (
		requestParam dtos.LoginRequestBody
	)

	if err := ctx.Bind(requestParam); err != nil {
		data := dtos.NewResponse(http.StatusUnprocessableEntity, "validations error", err, ctl.logger)
		return ctx.JSON(http.StatusUnprocessableEntity, data)
	}

	return ctx.NoContent(200)
}

func (ctl AuthControllerImpl) Register(ctx echo.Context) error {
	return ctx.String(200, "Register route")
}
