package sse

import (
	"fmt"
	"github.com/aasumitro/gowa/internal/delivery/middlewares"
	"github.com/gin-gonic/gin"
	"io"
	"log"
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

func (handler *StreamSSEHandler) qrStream(ctx *gin.Context) {
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
			log.Printf("Client added. %d registered clients", len(handler.TotalClients))

		// Remove closed client
		case client := <-handler.ClosedClients:
			delete(handler.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(handler.TotalClients))

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

func NewStreamSSEHandler(router *gin.RouterGroup) {
	handler := StreamSSEHandler{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	// We are streaming current time to clients in the interval 10 seconds
	go func() {
		for {
			time.Sleep(1 * time.Second)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)
			handler.Message <- currentTime
		}
	}()

	go handler.listen()

	router.GET("/stream", middlewares.SSEHeadersMiddleware(),
		handler.serve(), handler.qrStream)
}
