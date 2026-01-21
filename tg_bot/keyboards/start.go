package keyboards

import (
	"log"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/go-telegram/bot/models"
)

func StartKeyboard() *models.InlineKeyboardMarkup {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}

	var keyboard [][]models.InlineKeyboardButton

	keyboard = append(keyboard, []models.InlineKeyboardButton{
		{Text: "🔑 Ключ", CallbackData: "key"}})

	keyboard = append(keyboard, []models.InlineKeyboardButton{
		{Text: "ℹ️ Канал", URL: cfg.TelegramBot.Channel},
		{Text: "✉️ Поддержка", URL: cfg.TelegramBot.Support},
	})

	keyboard = append(keyboard, []models.InlineKeyboardButton{
		{Text: "📍 Дополнительное", CallbackData: "additionally"},
	})

	return &models.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}
