package handlers

import (
	httpDelivery "github.com/aasumitro/gowa/internal/delivery"
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// WhatsappMessageHandler struct
type whatsappMessageHTTPHandler struct {
	waService domain.WhatsappServiceContract
}

// NewWhatsappMessageHttpHandler constructor
// @params *gin.Engine
// @params domain.WhatsappServiceContract
func NewWhatsappMessageHttpHandler(
	router gin.IRoutes,
	waService domain.WhatsappServiceContract,
) {
	// Create a new handler and inject dependencies into it for use in the HTTP request handlers below.
	handler := &whatsappMessageHTTPHandler{waService: waService}

	// whatsapp message routes registration here ...
	router.POST("/send-text", handler.sendText)
	router.POST("/send-location", handler.sendLocation)
	router.POST("/send-image", handler.sendImage)
	router.POST("/send-audio", handler.sendAudio)
	router.POST("/send-document", handler.sendDocument)
}

// sendText handler
func (handler whatsappMessageHTTPHandler) sendText(context *gin.Context) {
	httpDelivery.NewHttpRespond(
		context,
		http.StatusOK,
		"Hello Text",
	)
}

// sendLocation handler
func (handler whatsappMessageHTTPHandler) sendLocation(context *gin.Context) {
	// TODO
}

// sendImage handler
func (handler whatsappMessageHTTPHandler) sendImage(context *gin.Context) {
	// TODO
}

// sendAudio handler
func (handler whatsappMessageHTTPHandler) sendAudio(context *gin.Context) {
	// TODO
}

// sendDocument handler
func (handler whatsappMessageHTTPHandler) sendDocument(context *gin.Context) {
	// TODO
}
