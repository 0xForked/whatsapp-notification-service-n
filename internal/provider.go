package internal

import (
	"context"
	sseHandler "github.com/aasumitro/gowa/internal/delivery/handlers/sse"
	"github.com/gin-gonic/gin"
)

func NewAPIProvider(ctx context.Context, router *gin.Engine) {
	v1 := router.Group("/api/v1")
	sseHandler.NewStreamSSEHandler(v1)
}
