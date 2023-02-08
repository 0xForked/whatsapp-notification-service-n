package rest

import (
	"fmt"
	"github.com/aasumitro/gowa/domain"
	"github.com/aasumitro/gowa/pkg/utils"
	"github.com/aasumitro/gowa/pkg/whatsapp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GowansRESTHandler struct {
	wa *whatsapp.Client
}

// Status godoc
// @Schemes
// @Summary 	 Account Status
// @Description  Get Current Account Authentication Status.
// @Tags 		 Whatsapp
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond "BASIC RESPOND"
// @Success 201 {object} utils.SuccessRespond "QR RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/whatsapp/status [GET]
func (handler *GowansRESTHandler) Status(ctx *gin.Context) {
	if handler.wa.WAC.Store.ID != nil && handler.wa.WAC.IsLoggedIn() {
		// replace * with <b> in frontend
		utils.NewHTTPRespond(ctx, http.StatusOK, fmt.Sprintf(
			"Logged in as *%s*",
			handler.wa.WAC.Store.PushName,
		))
		return
	}

	qrData, err := handler.wa.WAC.GetQRChannel(ctx)
	if err != nil {
		utils.NewHTTPRespond(ctx, http.StatusInternalServerError, fmt.Sprintf(
			"Someting wrong with the server, %s",
			err.Error(),
		))
		return
	}

	if err := handler.wa.WAC.Connect(); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusInternalServerError, fmt.Sprintf(
			"Someting wrong with the server, %s",
			err.Error(),
		))
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusCreated, <-qrData)
}

// Message godoc
// @Schemes
// @Summary 		Send Text Message
// @Description		Send Text Message to Specified Receipant
// @Tags 		 	Whatsapp
// @Accept 			mpfd
// @Produce 		json
// @Param 			msisdn formData string true "destination number. eg: 6281255423"
// @Param 			text formData string true "message text"
// @Success 200 {object} utils.SuccessRespond "BASIC RESPOND"
// @Success 422 {object} utils.ErrorRespond "UNPROCESSABLE ENTITY RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/whatsapp/message [POST]
func (handler *GowansRESTHandler) Message(ctx *gin.Context) {
	var form domain.TextMessage
	if err := ctx.ShouldBind(&form); err != nil {
		utils.NewHTTPRespond(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	id, err := handler.wa.SendMessage(ctx, form.Msisdn, form.Text)
	if err != nil {
		utils.NewHTTPRespond(ctx, http.StatusInternalServerError, fmt.Sprintf(
			"Someting wrong with the server, %s",
			err.Error(),
		))
		return
	}

	utils.NewHTTPRespond(ctx, http.StatusOK, id)
}

// Logout godoc
// @Schemes
// @Summary 	 Account Logout
// @Description  Logged Out Current Authenticate Account
// @Tags 		 Whatsapp
// @Accept       json
// @Produce      json
// @Success 401 {object} utils.ErrorRespond "UNAUTHORIZED RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/whatsapp/logout [POST]
func (handler *GowansRESTHandler) Logout(ctx *gin.Context) {
	err := handler.wa.WAC.Logout()

	code, data := func() (int, string) {
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		return http.StatusUnauthorized, "logged out"
	}()

	utils.NewHTTPRespond(ctx, code, data)
}

func NewGowansRESTHandler(router *gin.RouterGroup, wa *whatsapp.Client) {
	handler := GowansRESTHandler{wa: wa}
	router.GET("/whatsapp/status", handler.Status)

	protected := router.Use(func(ctx *gin.Context) {
		if !wa.WAC.IsLoggedIn() {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	})

	protected.POST("/whatsapp/message", handler.Message)
	protected.POST("/whatsapp/logout", handler.Logout)
}
