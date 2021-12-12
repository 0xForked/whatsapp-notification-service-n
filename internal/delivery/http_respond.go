package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HttpSuccessRespond message
type HttpSuccessRespond struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"OK"`
	Data   interface{} `json:"data"`
}

// HttpErrorRespond message
type HttpErrorRespond struct {
	Code   int    `json:"code" example:"400"`
	Status string `json:"status" example:"Bad Request"`
	Data   string `json:"data"`
}

// HttpValidationErrorRespond message
type HttpValidationErrorRespond struct {
	Code   int         `json:"code" example:"422"`
	Status string      `json:"status" example:"Unprocessable Entity"`
	Data   interface{} `json:"data"`
}

// HttpServerErrorRespond message
type HttpServerErrorRespond struct {
	Code   int    `json:"code" example:"500"`
	Status string `json:"status" example:"Internal Server Error"`
	Data   string `json:"data"`
}

// NewHttpRespond godoc
func NewHttpRespond(context *gin.Context, code int, data interface{}) {
	switch code {
	case http.StatusInternalServerError:
		context.JSON(
			code,
			HttpServerErrorRespond{
				Code:   code,
				Status: http.StatusText(code),
				Data:   "something went wrong with the server",
			},
		)
		break
	case http.StatusBadRequest:
		context.JSON(
			code,
			HttpErrorRespond{
				Code:   code,
				Status: http.StatusText(code),
				Data:   "something went wrong with the request",
			},
		)
		break
	case http.StatusUnprocessableEntity:
		context.JSON(
			code,
			HttpValidationErrorRespond{
				Code:   code,
				Status: http.StatusText(code),
				Data:   data,
			},
		)
		break
	default:
		context.JSON(
			code,
			HttpSuccessRespond{
				Code:   code,
				Status: http.StatusText(code),
				Data:   data,
			},
		)
		break
	}
}
