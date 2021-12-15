package handlers

import (
	httpDelivery "github.com/aasumitro/gowa/internal/delivery"
	"github.com/aasumitro/gowa/internal/domain/contracts"
	"github.com/aasumitro/gowa/internal/domain/models"
	"github.com/aasumitro/gowa/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// WhatsappMessageHandler struct
type whatsappMessageHTTPHandler struct {
	waService contracts.WhatsappService
}

// NewWhatsappMessageHttpHandler constructor
// @params *gin.Engine
// @params domain.WhatsappServiceContract
func NewWhatsappMessageHttpHandler(
	router gin.IRoutes,
	waService contracts.WhatsappService,
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
	var form models.WhatsappSendTextForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
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
	var form models.WhatsappSendLocationForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendImage handler
func (handler whatsappMessageHTTPHandler) sendImage(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendAudio handler
func (handler whatsappMessageHTTPHandler) sendAudio(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendDocument handler
func (handler whatsappMessageHTTPHandler) sendDocument(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}
