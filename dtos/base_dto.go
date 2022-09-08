package dtos

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/saintox/go-basic-auth/pkg/logger"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func NewResponse(code int, message string, data interface{}, logger *logger.Logger) (response Response) {
	if code >= 300 {
		var err error
		if err, ok := data.(*echo.HTTPError); ok {
			if code == 0 {
				code = err.Code
			}

			response = Response{
				Code:    code,
				Message: message,
				Error:   err.Message,
			}
		} else if err, ok := data.(error); ok {
			response = Response{
				Code:    code,
				Message: message,
				Error:   err.Error(),
			}
		} else {
			response = Response{
				Code:    code,
				Message: message,
				Error:   message,
			}
		}

		if logger != nil {
			if err != nil {
				logger.Error().Msg(err.Error())
			} else {
				logger.Error().Msg(errors.New(message).Error())
			}
		}

		return response
	} else {
		return Response{
			Code:    code,
			Message: message,
			Data:    data,
		}
	}
}
