package middlewares

import (
	httpDelivery "github.com/aasumitro/gowa/internal/delivery/http"
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (middleware HttpMiddleware) WhatsappSession(waService domain.WhatsappServiceContract) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := waService.HasSession(); err != nil {
			httpDelivery.NewHttpRespond(
				context,
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			)

			context.Abort()
		}

		context.Next()
	}
}
