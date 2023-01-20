package bot_test

import (
	"context"
	"github.com/aasumitro/gowa/pkg/bot"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChatBot_NewGPT3Bot(t *testing.T) {
	chatBot := bot.NewChatBot(
		bot.WithContext(context.TODO()),
		bot.WithAPIKey("lorem"))
	gptBot := chatBot.NewGPT3Bot()
	assert.NotNil(t, gptBot)
}
