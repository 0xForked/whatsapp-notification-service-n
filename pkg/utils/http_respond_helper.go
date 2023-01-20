package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessRespond message
type SuccessRespond struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// ErrorRespond message
type ErrorRespond struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

func NewHTTPRespond(context *gin.Context, code int, data interface{}) {
	if code == http.StatusOK || code == http.StatusCreated {
		context.JSON(code, SuccessRespond{
			Code:   code,
			Status: http.StatusText(code),
			Data:   data,
		})

		return
	}

	msg := func() string {
		switch {
		case data != nil:
			return data.(string)
		case code == http.StatusBadRequest:
			return "something went wrong with the request"
		default:
			return "something went wrong with the server"
		}
	}()

	context.JSON(code, ErrorRespond{
		Code:   code,
		Status: http.StatusText(code),
		Data:   msg,
	})
}
