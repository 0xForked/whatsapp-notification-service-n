package middlewares

import (
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/gin-gonic/gin"
)

// CORS will handle the CORS middleware
func (middleware HttpMiddleware) CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", domain.ORIGIN)
		context.Writer.Header().Add("Access-Control-Allow-Methods", domain.METHODS)
		context.Writer.Header().Add("Access-Control-Allow-Headers", domain.HEADERS)
		context.Next()
	}
}
