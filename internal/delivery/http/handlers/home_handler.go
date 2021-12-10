package handlers

import (
	httpDelivery "github.com/aasumitro/gowa/internal/delivery/http"
	"github.com/aasumitro/gowa/internal/domain"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type homeHTTPHandler struct{}

func NewHomeHttpHandler(
	router *gin.Engine,
) {
	handler := &homeHTTPHandler{}
	router.GET("/", handler.home)
	router.GET("/ping", handler.ping)
	router.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
	router.NoRoute(handler.notFound)
	router.NoMethod(handler.noMethod)
}

func (handler homeHTTPHandler) home(context *gin.Context) {
	httpDelivery.NewHttpRespond(
		context,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		map[string]interface{}{
			"title":       "Whatsapp API with Golang (GOWA)",
			"description": "GOWA is build on top of Gin Gonic Framework",
			"url": map[string]string{
				"docs": "http://localhost:8080/docs/index.html",
			},
		},
	)
}

func (handler homeHTTPHandler) ping(context *gin.Context) {
	httpDelivery.NewHttpRespond(
		context,
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"service is running well",
	)
}

func (handler homeHTTPHandler) notFound(context *gin.Context) {
	httpDelivery.NewHttpRespond(
		context,
		http.StatusNotFound,
		http.StatusText(http.StatusNotFound),
		domain.ErrRouteNotFound.Error(),
	)
}

func (handler homeHTTPHandler) noMethod(context *gin.Context) {
	httpDelivery.NewHttpRespond(
		context,
		http.StatusNotFound,
		http.StatusText(http.StatusNotFound),
		domain.ErrMethodNotFound.Error(),
	)
}
