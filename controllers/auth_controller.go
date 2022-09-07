package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/saintox/go-basic-auth/services"
)

type AuthController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

type AuthControllerImpl struct {
	//validator
	service *services.Service
}

func NewAuthController(service *services.Service) *AuthControllerImpl {
	return &AuthControllerImpl{
		service: service,
	}
}

func (ctl AuthControllerImpl) Login(ctx echo.Context) error {
	return ctx.String(200, "Login route")
}

func (ctl AuthControllerImpl) Register(ctx echo.Context) error {
	return ctx.String(200, "Register route")
}
