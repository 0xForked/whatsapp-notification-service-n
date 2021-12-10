package middlewares

import "github.com/gin-gonic/gin"

// Auth will handle the Auth middleware
func (middleware HttpMiddleware) Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// TODO
	}
}
