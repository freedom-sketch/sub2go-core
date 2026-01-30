package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/freedom-sketch/sub2go-core/tg_bot/keyboards"
	"github.com/freedom-sketch/sub2go-core/tg_bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	userName := update.Message.From.FirstName
	userUUID := utils.IntToUUID(update.Message.From.ID)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        fmt.Sprintf("%s, %s!", utils.Greeting(), userName),
		ReplyMarkup: keyboards.StartKeyboard(userUUID, cfg),
	})

	if err != nil {
		log.Panicf("Error sending start message: %v\n", err)
	}
}
