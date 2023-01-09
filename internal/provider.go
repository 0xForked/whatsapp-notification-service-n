package internal

import (
	"context"
	"fmt"
	"github.com/aasumitro/gowa/configs"
	"github.com/aasumitro/gowa/pkg/whatsapp"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

func NewAPIProvider(ctx context.Context, router *gin.Engine) {
	//_ = sqlRepository.NewSessionSQLRepository()
	wa := whatsapp.NewInstance(
		whatsapp.WithConfig(configs.Instance))
	client := wa.MakeConnection()
	client.AddEventHandler(func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			fmt.Println("msg", v.Message.GetConversation())
			fmt.Println("user", v.Info.Sender.User)
			if v.Message.GetConversation() == "BOT" && v.Info.Sender.User == "628817454334" {
				t, e := client.SendMessage(ctx, types.JID{
					User:   v.Info.Sender.User,
					Server: types.DefaultUserServer,
				}, &waProto.Message{
					Conversation: proto.String("ape lu kntl"),
				})
				if e != nil {
					fmt.Println(e.Error())
				}

				fmt.Println("SEND MESSAGE", t.ID)
			}
		}
	})
	//defer client.Disconnect()
	if client.Store.ID != nil {
		fmt.Println("data is not nil")
		err := client.Connect()
		if err != nil {
			fmt.Printf("[Main][Connect] Err: %+v\n", err)
			panic(err)
		}
	}

	//engine := router.Group("/api/v1")
	//engine.GET("test", func(c *gin.Context) {
	// No ID stored, new login
	//qrChan, _ := client.GetQRChannel(context.Background())
	//err := client.Connect()
	//if err != nil {
	//	fmt.Printf("[Main][GetQRChannel] Err: %+v\n", err)
	//	panic(err)
	//}
	//c.JSON(200, <-qrChan)
	//for evt := range qrChan {
	//	if evt.Event == "code" {
	//		c.JSON(200, evt.Code)
	//		//qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
	//	}
	//} else {
	//	fmt.Println("Login event:", evt.Event)
	//}
	//}
	//})
	//fmt.Println(ctx)
	//fmt.Println(engine)
}
