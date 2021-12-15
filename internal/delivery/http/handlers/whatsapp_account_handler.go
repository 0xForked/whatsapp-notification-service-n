package handlers

import (
	"github.com/aasumitro/gowa/internal/delivery"
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// whatsappAccountHTTPHandler struct
type whatsappAccountHTTPHandler struct {
	waService domain.WhatsappServiceContract
}

// NewWhatsappAccountHttpHandler constructor
// @params *gin.Engine
// @params domain.WhatsappServiceContract
func NewWhatsappAccountHttpHandler(
	router gin.IRoutes,
	waService domain.WhatsappServiceContract,
) {
	// Create a new handler and inject dependencies into it for use in the HTTP request handlers below.
	handler := &whatsappAccountHTTPHandler{waService: waService}

	// whatsapp message routes registration here ...
	router.POST("/login", handler.login)
	router.GET("/profile", handler.profile)
	router.POST("/logout", handler.logout)
}

func (handler whatsappAccountHTTPHandler) login(context *gin.Context) {
	// TODO: Implement the handler for the POST /login endpoint.
}

// profile godoc
// @Schemes
// @summary 	current connected account
// @Description Get logged in account profile
// @Tags 		Whatsapp
// @Accept  	json
// @Produce  	json
// @Success 200 {object} delivery.HttpSuccessRespond{data=object} "success respond"
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Failure 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/profile [get]
func (handler whatsappAccountHTTPHandler) profile(context *gin.Context) {
	// TODO: Implement the handler for the GET /profile endpoint.
}

// logout godoc.
// @Schemes
// @Summary 	Logout
// @Description Logout from whatsapp web.
// @Tags 		Whatsapp
// @Accept 		json
// @Produce 	json
// @Success 200 {object} delivery.HttpSuccessRespond{data=string} "success respond"
// @Success 400 {object} delivery.HttpErrorRespond{data=string} "bad request respond"
// @Success 500 {object} delivery.HttpServerErrorRespond{data=string} "internal server error respond"
// @Router /api/v1/whatsapp/logout [post]
func (handler whatsappAccountHTTPHandler) logout(context *gin.Context) {
	err := handler.waService.Logout()
	if err != nil {
		delivery.NewHttpRespond(
			context,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	delivery.NewHttpRespond(
		context,
		http.StatusOK,
		"[Action] Logout successfully",
	)
}
