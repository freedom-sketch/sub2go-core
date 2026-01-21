package handlers

import (
	"context"
	"fmt"

	"github.com/freedom-sketch/sub2go-core/tg_bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userName := update.Message.From.FirstName
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("%s, %s", utils.Greeting(), userName),
	})
}
