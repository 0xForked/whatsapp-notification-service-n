package whatsapp

import (
	"context"
	"fmt"
	"github.com/aasumitro/gowa/configs"
	"github.com/golang/protobuf/proto"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"sync"
)

var wacSingleton sync.Once

type (
	Client struct {
		WAC *whatsmeow.Client
	}

	IWhatsapp interface {
		MakeConnection() *Client
		SendMessage(ctx context.Context, recipient, message string) (string, error)
	}
)

func (c *Client) MakeConnection() *Client {
	container := sqlstore.NewWithDB(configs.DbPool, "sqlite3", nil)
	device, err := container.GetFirstDevice()
	if err != nil {
		panic(fmt.Sprintf("WHATSAPPMEOW_ERROR: %s", err.Error()))
	}

	wacSingleton.Do(func() {
		c.WAC = whatsmeow.NewClient(device, nil)
		if c.WAC.Store.ID != nil {
			if err := c.WAC.Connect(); err != nil {
				panic(fmt.Sprintf("WHATSAPPMEOW_ERROR: %s", err.Error()))
			}
		}
	})

	return c
}

func (c *Client) SendMessage(ctx context.Context, recipient, message string) (string, error) {
	resp, err := c.WAC.SendMessage(ctx, types.JID{
		User:   recipient,
		Server: types.DefaultUserServer,
	}, &waProto.Message{
		Conversation: proto.String(message),
	})

	return resp.ID, err
}
