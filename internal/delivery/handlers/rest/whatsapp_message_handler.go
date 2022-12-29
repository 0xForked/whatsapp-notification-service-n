package rest

import (
	"fmt"
	"github.com/aasumitro/gowa/internal/delivery/middlewares"
	"github.com/aasumitro/gowa/internal/domain/contracts"
	"github.com/aasumitro/gowa/internal/domain/models"
	"github.com/aasumitro/gowa/internal/utils"
	"github.com/aasumitro/gowa/pkg/apputils"
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

	sessionMiddleware := middlewares.WhatsappSession(waService)
	protectedRouter := router.Use(sessionMiddleware)
	// whatsapp message routes registration here ...
	protectedRouter.POST("/send-text", handler.sendText)
	protectedRouter.POST("/send-location", handler.sendLocation)
	protectedRouter.POST("/send-image", handler.sendImage)
	protectedRouter.POST("/send-audio", handler.sendAudio)
	protectedRouter.POST("/send-document", handler.sendDocument)
}

// sendText handler for send whatsapp message by text.
// @Schemes
// @Summary send text message
// @Description Send text via whatsapp message.
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
		apputils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	msgId, err := handler.waService.SendText(form)

	if err != nil {
		apputils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	apputils.NewHttpRespond(
		context,
		http.StatusOK,
		map[string]string{
			"message_id": msgId,
		},
	)
}

// sendLocation handler for send whatsapp message by location
// @Schemes
// @Summary send location message
// @Description send location via whatsapp message
// @Tags Whatsapp Messaging
// @Accept x-www-form-urlencoded
// @Produce json
// @Param msisdn formData string true "Destination number. eg: 6281255423 or group_creator-timstamp_created -> 6281271471566-1619679643 for group"
// @Param latitude formData number true "Latitude. e.g: 1.XXX"
// @Param longitude formData number true "Longitude. e.g: 124.XXX"
// @Success 200 {object} delivery.HttpSuccessRespond{data=object}
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 422 {object} delivery.HttpValidationErrorRespond{data=object} "unprocessable entity respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/send-location [POST]
func (handler whatsappMessageHTTPHandler) sendLocation(context *gin.Context) {
	var form models.WhatsappSendLocationForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		apputils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	msgId, err := handler.waService.SendLocation(form)

	if err != nil {
		apputils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	apputils.NewHttpRespond(
		context,
		http.StatusOK,
		map[string]string{
			"message_id": msgId,
		},
	)
}

// sendImage handler for send whatsapp message by image
// @Schemes
// @Summary send image message
// @Description send image via whatsapp message
// @Tags Whatsapp Messaging
// @Accept mpfd
// @Produce json
// @Param msisdn formData string true "Destination number. eg: 6281255423 or group_creator-timstamp_created -> 6281271471566-1619679643 for group"
// @Param file formData file true "Image. with extension: jpg,jpeg,png, with min length: 1024mb"
// @Success 200 {object} delivery.HttpSuccessRespond{data=object}
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 422 {object} delivery.HttpValidationErrorRespond{data=object} "unprocessable entity respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/send-image [POST]
func (handler whatsappMessageHTTPHandler) sendImage(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		apputils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	if form.FileHeader != nil {
		acceptedType := []string{"image/png", "image/jpg", "image/jpeg"}
		contentType := form.FileHeader.Header.Get("Content-Type")
		imageType := utils.Explode("/", contentType)

		if !utils.InArray(contentType, acceptedType) {
			apputils.NewHttpRespond(
				context,
				http.StatusBadRequest,
				fmt.Sprintf("file type error. accepted:png,jpg,jpeg. given:%v.", imageType[1]),
			)
			return
		}
	}

	msgId, err := handler.waService.SendFile(form, "image")

	if err != nil {
		apputils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	apputils.NewHttpRespond(
		context,
		http.StatusOK,
		map[string]string{
			"message_id": msgId,
		},
	)
}

// sendAudio handler for send whatsapp message by audio
// @Schemes
// @Summary send audio message
// @Description send audio via whatsapp message
// @Tags Whatsapp Messaging
// @Accept  x-www-form-urlencoded
// @Produce json
// @Param msisdn formData string true "Destination number. eg: 6281255423 or group_creator-timstamp_created -> 6281271471566-1619679643 for group"
// @Param file formData file true "Audio. with extension: mp3,aac,m4a,amr,opus, with min length: 1024mb"
// @Success 200 {object} delivery.HttpSuccessRespond{data=object}
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 422 {object} delivery.HttpValidationErrorRespond{data=object} "unprocessable entity respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/send-audio [POST]
func (handler whatsappMessageHTTPHandler) sendAudio(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		apputils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	if form.FileHeader != nil {
		acceptedType := []string{"audio/mpeg"}
		contentType := form.FileHeader.Header.Get("Content-Type")
		audioType := utils.Explode("/", contentType)

		if !utils.InArray(contentType, acceptedType) {
			apputils.NewHttpRespond(
				context,
				http.StatusBadRequest,
				fmt.Sprintf("file type error. accepted:mp3,aac,m4a,amr,opus. given:%v.", audioType[1]),
			)
			return
		}
	}

	msgId, err := handler.waService.SendFile(form, "audio")

	if err != nil {
		apputils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	apputils.NewHttpRespond(
		context,
		http.StatusOK,
		map[string]string{
			"message_id": msgId,
		},
	)
}

// sendDocument handler for send whatsapp message by document
// @Schemes
// @Summary send document message
// @Description send document via whatsapp message
// @Tags Whatsapp Messaging
// @Accept  x-www-form-urlencoded
// @Produce json
// @Param msisdn formData string true "Destination number. eg: 6281255423 or group_creator-timstamp_created -> 6281271471566-1619679643 for group"
// @Param file formData file true "Document. with extension: any, with min length: 1024mb"
// @Success 200 {object} delivery.HttpSuccessRespond{data=object}
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 422 {object} delivery.HttpValidationErrorRespond{data=object} "unprocessable entity respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/send-document [POST]
func (handler whatsappMessageHTTPHandler) sendDocument(context *gin.Context) {
	var form models.WhatsappSendFileForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewValidationErrors(models.WhatsappValidationErrorMessage).All(form, err)
		apputils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	msgId, err := handler.waService.SendFile(form, "document")

	if err != nil {
		apputils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	apputils.NewHttpRespond(
		context,
		http.StatusOK,
		map[string]string{
			"message_id": msgId,
		},
	)
}
