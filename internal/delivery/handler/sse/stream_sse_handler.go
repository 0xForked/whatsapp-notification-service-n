package sse

import (
	"context"
	"fmt"
	"github.com/aasumitro/gowa/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"io"
	"time"
)

type StreamSSEHandler struct {
	// Events are pushed to this channel by
	// the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	TotalClients map[chan string]bool
}

// ClientChan
// New event messages are broadcast to
// all registered client connection channels
type ClientChan chan string

func (handler *StreamSSEHandler) stream(ctx *gin.Context) {
	v, ok := ctx.Get("clientChan")
	if !ok {
		return
	}
	clientChan, ok := v.(ClientChan)
	if !ok {
		return
	}
	ctx.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		if msg, ok := <-clientChan; ok {
			ctx.SSEvent("message", msg)
			return true
		}
		return false
	})
}

// It Listens all incoming requests from clients.
// Handles addition and removal of clients and broadcast messages to clients.
func (handler *StreamSSEHandler) listen() {
	for {
		select {
		// Add new available client
		case client := <-handler.NewClients:
			handler.TotalClients[client] = true

		// Remove closed client
		case client := <-handler.ClosedClients:
			delete(handler.TotalClients, client)
			close(client)

		// Broadcast message to client
		case eventMsg := <-handler.Message:
			for clientMessageChan := range handler.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (handler *StreamSSEHandler) serve() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := make(ClientChan)

		// Send new connection to event server
		handler.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			handler.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}

func NewStreamSSEHandler(router *gin.RouterGroup, wac *whatsmeow.Client) {
	handler := StreamSSEHandler{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	// We are streaming current time to clients in the interval 10 seconds
	go func() {
		var msg string
		var nextCall time.Time
		data := func() string {
			if wac.Store.ID == nil {
				qrData, _ := wac.GetQRChannel(context.Background())

				if err := wac.Connect(); err != nil {
					return err.Error()
				}

				qr := <-qrData

				msg = qr.Code
				nextCall = time.Now().Add(60 * time.Second)
				return qr.Code
			}

			msg = ""
			nextCall = time.Now()
			return fmt.Sprintf(
				"You are logged in as %s",
				wac.Store.ID.User,
			)
		}

		wac.AddEventHandler(func(evt interface{}) {
			switch v := evt.(type) {
			case *events.Message:
				if !v.Info.IsFromMe {
					handler.Message <- fmt.Sprintf(
						"Log [%s] - %s (%s) send new message",
						v.Info.PushName, v.Info.Sender.User,
						time.Now().Format("2006-01-02 15:04:05"))
				}
			}
		})

		for {
			time.Sleep(1 * time.Second)

			if wac.Store.ID == nil && nextCall.After(time.Now()) {
				data()
			}

			if msg != "" {
				handler.Message <- msg
				continue
			}

			if wac.Store.ID != nil {
				handler.Message <- data()
			}
		}
	}()

	go handler.listen()

	router.GET("/stream",
		middleware.SSEHeadersMiddleware(),
		handler.serve(), handler.stream)
}
