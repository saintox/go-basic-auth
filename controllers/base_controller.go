package controllers

import "github.com/saintox/go-basic-auth/services"

type Controller struct {
	Auth AuthController
}

func NewController(service *services.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(service),
	}
}
