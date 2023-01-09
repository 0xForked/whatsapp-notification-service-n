package internal

import (
	"context"
	"fmt"
	"github.com/aasumitro/gowa/internal/delivery/handler/sse"
	"github.com/aasumitro/gowa/pkg/whatsapp"
	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow/types/events"
	"strings"
	"time"
)

var mainCtx context.Context
var waClient *whatsapp.Client

func NewAPIProvider(ctx context.Context, router *gin.Engine, whatsappClient *whatsapp.Client) {
	mainCtx, waClient = ctx, whatsappClient
	whatsappClient.WAC.AddEventHandler(stream)
	engine := router.Group("/api/v1")
	sse.NewStreamSSEHandler(engine, whatsappClient.WAC)
}

func stream(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		tmpl := strings.ToLower(v.Message.GetConversation())
		contains := strings.Contains(tmpl, "#bot")
		match := v.Info.Sender.User == "628817454334"
		if contains && match {
			time.AfterFunc(1*time.Second, func() {
				id, err := waClient.SendMessage(mainCtx, v.Info.Sender.User, "ape lu kntl")
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println("SEND MESSAGE", id)
			})
		}
	}
}
