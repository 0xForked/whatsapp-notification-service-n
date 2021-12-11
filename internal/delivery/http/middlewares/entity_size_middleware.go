package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

// EntitySizeAllowed will handle the incoming entity body size middleware
func (middleware HttpMiddleware) EntitySizeAllowed() gin.HandlerFunc {
	return func(context *gin.Context) {
		limit := os.Getenv("SERVER_UPLOAD_LIMIT")
		maxUploadSize, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		context.Request.Body = http.MaxBytesReader(
			context.Writer,
			context.Request.Body,
			maxUploadSize,
		)

		context.Next()
	}
}
