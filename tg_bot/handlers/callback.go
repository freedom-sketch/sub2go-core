package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/freedom-sketch/sub2go-core/infra/database"
	"github.com/freedom-sketch/sub2go-core/tg_bot/keyboards"
	"github.com/freedom-sketch/sub2go-core/tg_bot/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandleCallbackKey(ctx context.Context, b *bot.Bot, update *models.Update) {
	query := update.CallbackQuery
	if query == nil {
		return
	}

	cfg, err := config.Load("config.json")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}

	db, err := database.Connect(&cfg.DataBase)
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	userUUID := utils.IntToUUID(query.From.ID)

	subscription, err := database.GetSubscriptionByUserUUID(db, userUUID)
	if err != nil {
		log.Printf("Failed to get subscription: %v", err)
		return
	}

	subKey := utils.GenerateSubscriptionKey(userUUID)
	message := fmt.Sprintf("Ваш ключ: %s\nДата окончания: %s", subKey, subscription.EndDate.Format("2006-01-02"))

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{})

	editParams := &bot.EditMessageTextParams{
		ChatID:      query.From.ID,
		MessageID:   query.Message.Message.ID,
		Text:        message,
		ReplyMarkup: keyboards.StartKeyboard(userUUID),
	}
	b.EditMessageText(ctx, editParams)
}
