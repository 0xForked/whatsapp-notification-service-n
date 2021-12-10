package handlers

import (
	"github.com/gin-gonic/gin"
)

type whatsappMessageHTTPHandler struct{}

func NewWhatsappMessageHttpHandler(
	router *gin.Engine,
) {
	//handler := &whatsappMessageHTTPHandler{}
	//v1 := router.Group("/api/v1")
	//v1.GET("/api", handler.home)
}

func (handler whatsappMessageHTTPHandler) SendText(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) SendLocation(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) SendImage(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) SendAudio(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) SendDocument(context *gin.Context) {
	// TODO
}

func (handler whatsappMessageHTTPHandler) Groups(context *gin.Context) {
	// TODO
}
