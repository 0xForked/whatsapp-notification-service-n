package handlers

import (
	"fmt"
	"github.com/aasumitro/gowa/internal/domain/contracts"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// ExampleHandler represent the http handler for example
type whatsappLoginWSHandler struct {
	waService  contracts.WhatsappService
	wsUpgrader websocket.Upgrader
}

func NewWhatsappLoginWSHandler(
	router *gin.Engine,
	waService contracts.WhatsappService,
) {
	handler := &whatsappLoginWSHandler{
		wsUpgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		waService: waService,
	}
	router.GET("/login", handler.login)
}

func (handler *whatsappLoginWSHandler) login(context *gin.Context) {
	conn, err := handler.wsUpgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	var qrCodeStr string
	done := make(chan bool)
	ticker := time.NewTicker(5 * time.Second)

	for {
		qrCodeStr = "1234"

		go writer(conn, []byte(qrCodeStr))

		for {
			select {
			case <-done:
				go writer(conn, []byte("done and close channel"))
				ticker.Stop()
				return
			case t := <-ticker.C:
				qrCodeStr = t.String()

				go writer(conn, []byte(qrCodeStr))

				time.Sleep(10 * time.Second)
				done <- true
			}
		}
	}
}

func login() {

}

func writer(conn *websocket.Conn, msg []byte) {
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		fmt.Printf("Failed to write message: %+v\n", err)
	}
}
