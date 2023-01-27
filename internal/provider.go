package internal

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/aasumitro/gowa/configs"
	"github.com/aasumitro/gowa/constants"
	"github.com/aasumitro/gowa/internal/delivery/handler/rest"
	"github.com/aasumitro/gowa/internal/delivery/handler/ws"
	"github.com/aasumitro/gowa/pkg/bot"
	"github.com/aasumitro/gowa/pkg/utils"
	"github.com/aasumitro/gowa/pkg/whatsapp"
	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow/types/events"
	"strings"
	"time"
)

func NewAPIProvider(ctx context.Context, router *gin.Engine, whatsappClient *whatsapp.Client) {
	chatBot := bot.NewChatBot(bot.WithContext(ctx),
		bot.WithAPIKey(configs.Instance.Gpt3APIKey))
	gptBot := chatBot.NewGPT3Bot()

	whatsappClient.WAC.AddEventHandler(func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Disconnected:
			_ = whatsappClient.WAC.Connect()
		case *events.Message:
			var data []string
			tmpl := strings.ToLower(v.Message.GetConversation())
			contains := utils.InArray(strings.Trim(tmpl[:4], " "),
				[]string{"#bot", "@bot", "bot"})
			fmt.Println(contains)
			match := v.Info.Sender.User == configs.Instance.BotReqMSISDN
			if contains && match {
				if err := gptBot.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
					Prompt:      []string{tmpl[4:]},
					MaxTokens:   gpt3.IntPtr(constants.GptMaxToken),
					Temperature: gpt3.Float32Ptr(0),
				}, func(response *gpt3.CompletionResponse) {
					data = append(data, response.Choices[0].Text)

					if response.Choices[0].FinishReason == "stop" {
						_, _ = whatsappClient.SendMessage(ctx,
							v.Info.Sender.User, strings.Join(data[2:], ""))
					}

					fmt.Println(data)
				}); err == nil {
					_, _ = whatsappClient.SendMessage(ctx,
						v.Info.Sender.User, fmt.Sprintf(
							"i cant proceed your request, reason: %s",
							err.Error()))
				}
			}

			if v.Info.Sender.User != configs.Instance.BotReqMSISDN && !v.Info.IsFromMe {
				_, _ = whatsappClient.SendMessage(context.Background(),
					configs.Instance.BotReqMSISDN, fmt.Sprintf(
						"[%s] new message from %s (%s) - [%s]",
						time.Now().Format("2006-01-02 15:04:05"),
						v.Info.PushName, v.Info.Sender.User,
						v.Message.GetConversation()))

				time.Sleep(1 * time.Second)
			}
		}
	})

	engine := router.Group("/api/v1")
	ws.NewGowansWSHandler(engine, whatsappClient)
	rest.NewGowansRESTHandler(engine, whatsappClient)
}
