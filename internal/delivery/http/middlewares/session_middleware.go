package middlewares

import (
	httpDelivery "github.com/aasumitro/gowa/internal/delivery"
	"github.com/aasumitro/gowa/internal/domain/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (middleware HttpMiddleware) WhatsappSession(waService contracts.WhatsappService) gin.HandlerFunc {
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
