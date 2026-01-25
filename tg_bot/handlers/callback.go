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

func Key(ctx context.Context, b *bot.Bot, update *models.Update) {
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
	message := fmt.Sprintf(`🔑 Твой ключ: <code>%s</code>
⏳ Дата окончания: %s
📡 Трафик: %s`,
		subKey, subscription.EndDate.Format("2006.01.02"), utils.TrafficFormat(subscription))

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{})

	buttonBack := keyboards.ButtonBack()
	keyboard := [][]models.InlineKeyboardButton{buttonBack}

	editParams := &bot.EditMessageTextParams{
		ChatID:      query.From.ID,
		MessageID:   query.Message.Message.ID,
		Text:        message,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboard},
	}

	_, err = b.EditMessageText(ctx, editParams)
	if err != nil {
		log.Printf("Failed to edit message: %v", err)
	}
}

func Back(ctx context.Context, b *bot.Bot, update *models.Update) {
	query := update.CallbackQuery
	if query == nil {
		return
	}

	userName := query.From.FirstName
	userID := query.From.ID
	userUUID := utils.IntToUUID(userID)

	editParams := &bot.EditMessageTextParams{
		ChatID:      query.From.ID,
		MessageID:   query.Message.Message.ID,
		Text:        fmt.Sprintf("%s, %s", utils.Greeting(), userName),
		ReplyMarkup: keyboards.StartKeyboard(userUUID),
	}

	_, err := b.EditMessageText(ctx, editParams)
	if err != nil {
		log.Printf("Failed to edit message: %v", err)
	}
}
