package http

import "github.com/gin-gonic/gin"

// Respond message
type Respond struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// NewHttpRespond NewHttpErrorRespond NewHttpError
func NewHttpRespond(context *gin.Context, code int, status string, data interface{}) {
	context.JSON(
		code,
		Respond{
			Code:   code,
			Status: status,
			Data:   data,
		},
	)
}
