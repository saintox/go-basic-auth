package controllers

import (
	"github.com/saintox/go-basic-auth/middlewares"
	"github.com/saintox/go-basic-auth/pkg/logger"
	"github.com/saintox/go-basic-auth/services"
)

type Controller struct {
	Auth AuthController
}

func NewController(service *services.Service, validator *middlewares.Validator, logger *logger.Logger) *Controller {
	return &Controller{
		Auth: NewAuthController(service, validator, logger),
	}
}
