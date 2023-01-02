package internal

import (
	"context"
	"github.com/gin-gonic/gin"
)

func NewAPIProvider(ctx context.Context, router *gin.Engine) {
	_ = router.Group("/api/v1")
}
