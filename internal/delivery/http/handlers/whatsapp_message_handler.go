package handlers

import (
	httpDelivery "github.com/aasumitro/gowa/internal/delivery/http"
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type whatsappMessageHTTPHandler struct {
	waService domain.WhatsappServiceContract
}

func NewWhatsappMessageHttpHandler(
	router *gin.Engine,
	waService domain.WhatsappServiceContract,
) {
	handler := &whatsappMessageHTTPHandler{waService: waService}
	v1 := router.Group("/api/v1/whatsapp")
	v1.GET("/send-text", handler.sendText)
	v1.GET("/send-location", handler.sendLocation)
	v1.GET("/send-image", handler.sendImage)
	v1.GET("/send-audio", handler.sendAudio)
	v1.GET("/send-document", handler.sendDocument)
}

func (handler whatsappMessageHTTPHandler) sendText(context *gin.Context) {
	httpDelivery.NewHttpRespond(
		context,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"Hello Text",
	)
}

func (handler whatsappMessageHTTPHandler) sendLocation(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) sendImage(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) sendAudio(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) sendDocument(context *gin.Context) {
	// TODO
}
