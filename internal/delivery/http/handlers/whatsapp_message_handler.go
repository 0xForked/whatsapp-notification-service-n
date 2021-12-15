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
	var form domain.WhatsappSendTextForm

	if err := context.ShouldBind(&form); err != nil {
		ginError := domain.NewGinErrors(domain.RequestWhatsappErrorMessage)
		errs := ginError.ListAllErrors(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, errs)
		return
	}

	httpDelivery.NewHttpRespond(
		context,
		http.StatusOK,
		&form,
	)
}

// sendLocation handler
func (handler whatsappMessageHTTPHandler) sendLocation(context *gin.Context) {
	var form domain.WhatsappSendLocationForm

	if err := context.ShouldBind(&form); err != nil {
		ginError := domain.NewGinErrors(domain.RequestWhatsappErrorMessage)
		errs := ginError.ListAllErrors(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, errs)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendImage handler
func (handler whatsappMessageHTTPHandler) sendImage(context *gin.Context) {
	var form domain.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		ginError := domain.NewGinErrors(domain.RequestWhatsappErrorMessage)
		errs := ginError.ListAllErrors(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, errs)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendAudio handler
func (handler whatsappMessageHTTPHandler) sendAudio(context *gin.Context) {
	var form domain.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		ginError := domain.NewGinErrors(domain.RequestWhatsappErrorMessage)
		errs := ginError.ListAllErrors(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, errs)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendDocument handler
func (handler whatsappMessageHTTPHandler) sendDocument(context *gin.Context) {
	var form domain.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		ginError := domain.NewGinErrors(domain.RequestWhatsappErrorMessage)
		errs := ginError.ListAllErrors(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, errs)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}
