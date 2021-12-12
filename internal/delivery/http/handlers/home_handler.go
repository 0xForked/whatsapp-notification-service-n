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
	delivery.NewHttpRespond(
		context,
		http.StatusOK,
		map[string]interface{}{
			"01_title": fmt.Sprintf("Whatsapp API with Golang (%s)", os.Getenv("SERVER_SHORT_NAME")),
			"02_internal_link": map[string]string{
				"docs": fmt.Sprintf("http://%s/docs/index.html", os.Getenv("SERVER_URL")),
			},
			"03_perquisites": map[string]interface{}{
				"01_language": map[string]string{
					"home_page":  "https://golang.org/",
					"repository": "https://github.com/golang/go",
					"license":    "https://github.com/golang/go/blob/master/LICENSE",
				},
				"02_gin": map[string]string{
					"home_page":  "https://gin-gonic.com/",
					"repository": "https://github.com/gin-gonic/gin",
					"license":    "https://github.com/gin-gonic/gin/blob/master/LICENSE",
				},
				"03_gorilla_websocket": map[string]string{
					"repository": "github.com/gorilla/websocket",
					"license":    "https://github.com/gorilla/websocket/blob/master/LICENSE",
				},
				"04_swaggo": map[string]string{
					"swaggo_swag":  "github.com/swaggo/swag",
					"swaggo_gin":   "github.com/swaggo/gin-swagger",
					"swaggo_files": "github.com/swaggo/files",
					"license":      "https://github.com/swaggo/swag/blob/master/license",
				},
				"05_go_whatsapp": map[string]string{
					"repository": "https://github.com/Rhymen/go-whatsapp",
					"license":    "https://github.com/Rhymen/go-whatsapp/blob/master/LICENSE",
				},
			},
		})
}

func (handler homeHTTPHandler) ping(context *gin.Context) {
	delivery.NewHttpRespond(
		context,
		http.StatusOK,
		"PONG! {service is running well}",
	)
}

func (handler homeHTTPHandler) notFound(context *gin.Context) {
	delivery.NewHttpRespond(
		context,
		http.StatusNotFound,
		domain.ErrRouteNotFound.Error(),
	)
}

func (handler homeHTTPHandler) noMethod(context *gin.Context) {
	delivery.NewHttpRespond(
		context,
		http.StatusNotFound,
		domain.ErrMethodNotFound.Error(),
	)
}
