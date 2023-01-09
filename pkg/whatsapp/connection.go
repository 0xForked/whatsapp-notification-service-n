package whatsapp

import (
	"fmt"
	"github.com/aasumitro/gowa/configs"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type (
	client struct {
		config *configs.AppConfig
	}

	Option func(*client)

	Whatsapp interface {
		MakeConnection() *whatsmeow.Client
	}
)

func WithConfig(config *configs.AppConfig) Option {
	return func(c *client) {
		c.config = config
	}
}

func NewInstance(opts ...Option) Whatsapp {
	c := &client{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *client) MakeConnection() *whatsmeow.Client {
	dbLog := waLog.Stdout("DATABASE", "DEBUG", true)
	container, err := sqlstore.New(c.config.DBDriver, c.config.DBDsnURL, dbLog)
	if err != nil {
		panic(fmt.Sprintf("DATABASE_ERROR: %s", err.Error()))
	}

	device, err := container.GetFirstDevice()
	if err != nil {
		panic(fmt.Sprintf("WHATSAPPMEOW_ERROR: %s", err.Error()))
	}

	clientLog := waLog.Stdout("CLIENT", "DEBUG", true)
	return whatsmeow.NewClient(device, clientLog)
}
