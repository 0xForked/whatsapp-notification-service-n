package apputils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPSuccessRespond struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"OK"`
	Data   interface{} `json:"data"`
}

type HTTPErrorRespond struct {
	Code   int    `json:"code" example:"400"`
	Status string `json:"status" example:"Bad Request"`
	Data   string `json:"data"`
}

type HTTPValidationErrorRespond struct {
	Code   int         `json:"code" example:"422"`
	Status string      `json:"status" example:"Unprocessable Entity"`
	Data   interface{} `json:"data"`
}

type HTTPServerErrorRespond struct {
	Code   int    `json:"code" example:"500"`
	Status string `json:"status" example:"Internal Server Error"`
	Data   string `json:"data"`
}

func NewHTTPRespond(context *gin.Context, code int, data interface{}) {
	if code == http.StatusOK || code == http.StatusCreated {
		context.JSON(
			code,
			HTTPSuccessRespond{
				Code:   code,
				Status: http.StatusText(code),
				Data:   data,
			},
		)

		return
	}

	if code == http.StatusUnprocessableEntity {
		context.JSON(
			code,
			HTTPValidationErrorRespond{
				Code:   code,
				Status: http.StatusText(code),
				Data:   data,
			},
		)

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

	context.JSON(
		code,
		HTTPErrorRespond{
			Code:   code,
			Status: http.StatusText(code),
			Data:   msg,
		},
	)
}
