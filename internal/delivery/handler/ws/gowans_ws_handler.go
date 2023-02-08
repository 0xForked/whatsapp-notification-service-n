package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aasumitro/gowa/configs"
	"github.com/aasumitro/gowa/constants"
	"github.com/aasumitro/gowa/pkg/whatsapp"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mau.fi/whatsmeow/types/events"
	"sync"
	"time"
)

type GowansWSHandler struct {
	wsc *websocket.Conn
	wa  *whatsapp.Client
}

var (
	wsUpgraded = websocket.Upgrader{
		ReadBufferSize:  constants.MaxWSBufferSize,
		WriteBufferSize: constants.MaxWSBufferSize,
	}
	mu sync.Mutex
)

func (handler *GowansWSHandler) status() {
	if handler.wa.WAC.Store.ID == nil && !handler.wa.WAC.IsLoggedIn() {
		if qrData, err := handler.wa.WAC.GetQRChannel(context.Background()); err == nil {
			if err := handler.wa.WAC.Connect(); err != nil {
				handler.sendMessage("error", "status", "")
				return
			}

			for qr := range qrData {
				handler.sendMessage("success", "qrcode", qr.Code)
				if qr.Event == "success" {
					break
				}
				time.Sleep(1 * time.Second)
			}
		}
	}

	handler.sendMessage("success", "loggedIn", fmt.Sprintf(
		"Logged in as <b>%s</b>",
		handler.wa.WAC.Store.PushName,
	))
}

func (handler *GowansWSHandler) logout() {
	status, data := "error", "you are not logged in!"

	if handler.wa.WAC.IsLoggedIn() {
		err := handler.wa.WAC.Logout()

		status, data = func() (string, string) {
			if err != nil {
				return "error", err.Error()
			}
			return "success", ""
		}()
	}

	handler.sendMessage(status, "loggedOut", data)
}

func (handler *GowansWSHandler) sendMessage(status, action string, data any) {
	mu.Lock()
	msg, _ := json.Marshal(map[string]any{
		"status": status,
		"action": action,
		"data":   data,
	})
	_ = handler.wsc.WriteMessage(
		websocket.TextMessage, msg)
	mu.Unlock()
}

func (handler *GowansWSHandler) waEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Disconnected:
		_ = handler.wa.WAC.Connect()
	case *events.Message:
		if !v.Info.IsFromMe {
			handler.sendMessage("success", "log", fmt.Sprintf(
				"Log [%s] - %s (%s) send new message",
				time.Now().Format("2006-01-02 15:04:05"),
				v.Info.PushName, v.Info.Sender.User))

			if v.Info.Sender.User != configs.Instance.BotReqMSISDN {
				_, _ = handler.wa.SendMessage(context.Background(),
					configs.Instance.BotReqMSISDN, fmt.Sprintf(
						"[%s] new message from %s (%s) - [%s]",
						time.Now().Format("2006-01-02 15:04:05"),
						v.Info.PushName, v.Info.Sender.User,
						v.Message.GetConversation()))

				time.Sleep(1 * time.Second)
			}
		}
	}
}

func (handler *GowansWSHandler) wsEventHandler() {
	for {
		_, msg, err := handler.wsc.ReadMessage()
		if err != nil {
			break
		}

		switch string(msg) {
		case "status":
			handler.status()
		case "logout":
			handler.logout()
		}
	}
}

func (handler *GowansWSHandler) Serve(ctx *gin.Context) {
	handler.wsc, _ = wsUpgraded.Upgrade(ctx.Writer, ctx.Request, nil)
	defer func(ws *websocket.Conn) { _ = ws.Close() }(handler.wsc)
	handler.wa.WAC.AddEventHandler(
		handler.waEventHandler)
	handler.wsEventHandler()
}

func NewGowansWSHandler(router *gin.RouterGroup, wa *whatsapp.Client) {
	handler := GowansWSHandler{wa: wa}
	router.GET("/ws", handler.Serve)
}
