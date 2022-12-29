package middlewares

import (
	"github.com/aasumitro/gowa/internal/domain/contracts"
	httpDelivery "github.com/aasumitro/gowa/pkg/apputils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WhatsappSession(waService contracts.WhatsappService) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := waService.HasSession(); err != nil {
			httpDelivery.NewHttpRespond(
				context,
				http.StatusBadRequest,
				err.Error(),
			)

			context.Abort()
		}

		context.Next()
	}
}
