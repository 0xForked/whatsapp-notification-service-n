package handlers

import (
	"fmt"
	"github.com/aasumitro/gowa/internal/delivery"
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

type homeHTTPHandler struct{}

func NewHomeHttpHandler(
	router *gin.Engine,
) {
	handler := &homeHTTPHandler{}
	router.NoMethod(handler.noMethod)
	router.NoRoute(handler.notFound)
	router.GET("/", handler.home)
	router.GET("/ping", handler.ping)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (handler homeHTTPHandler) home(context *gin.Context) {
	delivery.NewHttpRespond(context, http.StatusOK, map[string]interface{}{
		"01_title": fmt.Sprintf("Whatsapp API with Golang (%s)", os.Getenv("SERVER_NAME")),
		"02_spec":  fmt.Sprintf("http://%s/docs/index.html", os.Getenv("SERVER_URL")),
		"03_perquisites": map[string]interface{}{
			"01_language":  "https://github.com/golang/go",
			"02_framework": "https://github.com/gin-gonic/gin",
			"03_library": map[string]string{
				"01_swagger":  "https://github.com/swaggo/swag",
				"02_whatsapp": "https://github.com/Rhymen/go-whatsapp",
			},
		},
	})
}

func (handler homeHTTPHandler) ping(context *gin.Context) {
	delivery.NewHttpRespond(context, http.StatusOK, domain.PONG)
}

func (handler homeHTTPHandler) notFound(context *gin.Context) {
	delivery.NewHttpRespond(context, http.StatusNotFound, domain.ErrRouteNotFound.Error())
}

func (handler homeHTTPHandler) noMethod(context *gin.Context) {
	delivery.NewHttpRespond(context, http.StatusNotFound, domain.ErrMethodNotFound.Error())
}
