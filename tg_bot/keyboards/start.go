package keyboards

import (
	"log"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/freedom-sketch/sub2go-core/infra/database"
	"github.com/go-telegram/bot/models"
	"github.com/google/uuid"
)

func StartKeyboard(UserUUID uuid.UUID) *models.InlineKeyboardMarkup {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	db, err := database.Connect(&cfg.DataBase)
	if err != nil {
		log.Panicf("failed to connect to database: %v", err)
	}

	var keyboard [][]models.InlineKeyboardButton

	isAdmin, err := database.IsAdmin(db, UserUUID)
	if err != nil {
		log.Panicf("error checking for administrator rights")
	}

	if isAdmin {
		keyboard = append(keyboard, []models.InlineKeyboardButton{
			{Text: "💻 Панель администратора", CallbackData: "admin_panel"}})
	}

	activeSub, err := database.HasActiveSubscription(db, UserUUID)
	if err != nil {
		log.Panicf("failed to check active subscription: %v", err)
	}

	if activeSub {
		keyboard = append(keyboard, []models.InlineKeyboardButton{
			{Text: "🔑 Мой ключ", CallbackData: "key"}})
	}

	keyboard = append(keyboard, []models.InlineKeyboardButton{
		{Text: "ℹ️ Канал", URL: cfg.TelegramBot.Channel},
		{Text: "✉️ Поддержка", URL: cfg.TelegramBot.Support},
	})

	keyboard = append(keyboard, []models.InlineKeyboardButton{
		{Text: "📍 Дополнительное", CallbackData: "additionally"},
	})

	return &models.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}
