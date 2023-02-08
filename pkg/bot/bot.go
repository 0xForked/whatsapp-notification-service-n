package bot

import (
	"context"
	"github.com/PullRequestInc/go-gpt3"
)

type ChatBot struct {
	ctx    context.Context
	apiKey string
}

type ChatBotOptions func(bot *ChatBot)

func WithContext(ctx context.Context) ChatBotOptions {
	return func(chatBot *ChatBot) {
		chatBot.ctx = ctx
	}
}

func WithAPIKey(apiKey string) ChatBotOptions {
	return func(chatBot *ChatBot) {
		chatBot.apiKey = apiKey
	}
}

func NewChatBot(opts ...ChatBotOptions) *ChatBot {
	chatBot := &ChatBot{}

	for _, opt := range opts {
		opt(chatBot)
	}

	return chatBot
}

func (bot *ChatBot) NewGPT3Bot() gpt3.Client {
	return gpt3.NewClient(bot.apiKey)
}
