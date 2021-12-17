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

// SendText func for send text.
// @Schemes
// @Summary send text message
// @Description Send whatsapp text message.
// @Tags Whatsapp Messaging
// @Accept mpfd
// @Produce json
// @Param msisdn formData string true "Destination number. eg: 6281255423 or group_creator-timstamp_created -> 6281271471566-1619679643 for group"
// @Param text formData string true "Message text"
// @Success 200 {object} delivery.HttpSuccessRespond{data=object} "success respond"
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 422 {object} delivery.HttpValidationErrorRespond{data=object} "unprocessable entity respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/send-text [POST]
func (handler whatsappMessageHTTPHandler) sendText(context *gin.Context) {
	var form models.WhatsappSendTextForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	msgId, err := handler.waService.SendText(form)
	if err != nil {
		httpDelivery.NewHttpRespond(context, http.StatusBadRequest, err.Error())
	}

	httpDelivery.NewHttpRespond(
		context,
		http.StatusOK,
		map[string]string{
			"message_id": msgId,
		},
	)
}

// sendLocation handler TODO
func (handler whatsappMessageHTTPHandler) sendLocation(context *gin.Context) {
	var form models.WhatsappSendLocationForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendImage handler TODO
func (handler whatsappMessageHTTPHandler) sendImage(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendAudio handler TODO
func (handler whatsappMessageHTTPHandler) sendAudio(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}

// sendDocument handler TODO
func (handler whatsappMessageHTTPHandler) sendDocument(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		httpDelivery.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	httpDelivery.NewHttpRespond(context, http.StatusOK, &form)
}
