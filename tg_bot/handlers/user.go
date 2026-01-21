package handlers

import (
	"context"
	"fmt"

	"github.com/freedom-sketch/sub2go-core/tg_bot/keyboards"
	"github.com/freedom-sketch/sub2go-core/tg_bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("StartHandler called")
	userName := update.Message.From.FirstName
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        fmt.Sprintf("%s, %s", utils.Greeting(), userName),
		ReplyMarkup: keyboards.StartKeyboard(),
	})

	if err != nil {
		fmt.Printf("Error sending start message: %v\n", err)
	}
}
