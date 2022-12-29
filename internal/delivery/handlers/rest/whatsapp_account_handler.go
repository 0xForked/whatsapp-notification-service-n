package rest

import (
	"github.com/aasumitro/gowa/internal/delivery/middlewares"
	"github.com/aasumitro/gowa/internal/domain/contracts"
	"github.com/aasumitro/gowa/pkg/apperrors"
	"github.com/aasumitro/gowa/pkg/apputils"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)

// whatsappAccountHTTPHandler struct
type whatsappAccountHTTPHandler struct {
	waService contracts.WhatsappService
}

// NewWhatsappAccountHttpHandler constructor
// @params *gin.Engine
// @params domain.WhatsappServiceContract
func NewWhatsappAccountHttpHandler(
	router gin.IRoutes,
	waService contracts.WhatsappService,
) {
	// Create a new handler and inject dependencies into it for use in the HTTP request handlers below.
	handler := &whatsappAccountHTTPHandler{waService: waService}

	// whatsapp message routes registration here ...
	router.POST("/login", handler.login)
	sessionMiddleware := middlewares.WhatsappSession(waService)
	protectedRouter := router.Use(sessionMiddleware)
	protectedRouter.GET("/profile", handler.profile)
	protectedRouter.POST("/logout", handler.logout)
}

// login godoc
// @Schemes
// @summary 	login handler
// @Description Get logged in to account
// @Tags 		Whatsapp Account
// @Produce  	json
// @Produce 	html
// @Success 201 {object} delivery.HttpSuccessRespond{data=object} "success respond application/json"
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/login [POST]
func (handler whatsappAccountHTTPHandler) login(context *gin.Context) {
	qrCodeStr, err := handler.waService.Login()

	if err != nil {
		if err.Error() == apperrors.ErrAlreadyConnectedAndLoggedIn.Error() {
			profile, _ := handler.waService.Profile()
			apputils.NewHttpRespond(context, http.StatusOK, profile)
			return
		}

		apputils.NewHttpRespond(
			context,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	if context.Request.Header.Get("accept") == "text/html" {
		qrCodePng, _ := qrcode.Encode(qrCodeStr, qrcode.Medium, 256)
		context.Set("content-type", "image/png")
		_, _ = context.Writer.Write(qrCodePng)
		return
	}

	apputils.NewHttpRespond(
		context,
		http.StatusCreated,
		map[string]string{
			"qrcode":       qrCodeStr,
			"refresh_time": "20",
			"refresh_mode": "SECOND",
		},
	)
}

// profile godoc
// @Schemes
// @summary 	current connected account
// @Description Get logged in account profile
// @Tags 		Whatsapp Account
// @Accept  	json
// @Produce  	json
// @Success 200 {object} delivery.HttpSuccessRespond{data=object} "success respond"
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/profile [get]
func (handler whatsappAccountHTTPHandler) profile(context *gin.Context) {
	profile, err := handler.waService.Profile()

	if err != nil {
		apputils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	apputils.NewHttpRespond(context, http.StatusOK, profile)
}

// logout godoc.
// @Schemes
// @Summary 	Logout
// @Description Logout from whatsapp web.
// @Tags 		Whatsapp Account
// @Accept 		json
// @Produce 	json
// @Success 200 {object} delivery.HttpSuccessRespond{data=string} "success respond"
// @Success 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Success 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/logout [post]
func (handler whatsappAccountHTTPHandler) logout(context *gin.Context) {
	err := handler.waService.Logout()

	if err != nil {
		apputils.NewHttpRespond(
			context,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	apputils.NewHttpRespond(
		context,
		http.StatusOK,
		"[Action] Logout successfully",
	)
}
