package middlewares

import (
	"github.com/aasumitro/gowa/pkg/appconsts"
	"github.com/gin-gonic/gin"
)

// CORS will handle the CORS middleware
func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", appconsts.Origin)
		context.Writer.Header().Add("Access-Control-Allow-Methods", appconsts.Methods)
		context.Writer.Header().Add("Access-Control-Allow-Headers", appconsts.Headers)
		context.Next()
	}
}
