package handlers

import (
	httpDelivery "github.com/aasumitro/gowa/internal/delivery/http"
	"github.com/aasumitro/gowa/internal/delivery/http/middlewares"
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
	router *gin.Engine,
	waService domain.WhatsappServiceContract,
) {
	// Create a new handler and inject dependencies into it for use in the HTTP request handlers below.
	handler := &whatsappMessageHTTPHandler{waService: waService}

	// register custom middleware
	httpMiddleware := middlewares.InitHttpMiddleware()
	// create a new router group for the handler to register routes to and apply the middleware to it.
	// The middleware will be applied to all the routes registered in this group.
	v1 := router.Group("/api/v1/whatsapp").Use(
		//	httpMiddleware.CORS(),
		//	httpMiddleware.EntitySizeAllowed(),
		httpMiddleware.WhatsappSession(handler.waService),
	)

	// whatsapp message routes registration here ...
	v1.GET("/send-text", handler.sendText)
	v1.GET("/send-location", handler.sendLocation)
	v1.GET("/send-image", handler.sendImage)
	v1.GET("/send-audio", handler.sendAudio)
	v1.GET("/send-document", handler.sendDocument)
}

// sendText handler
func (handler whatsappMessageHTTPHandler) sendText(context *gin.Context) {
	httpDelivery.NewHttpRespond(
		context,
		http.StatusOK,
		http.StatusText(http.StatusOK),
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
