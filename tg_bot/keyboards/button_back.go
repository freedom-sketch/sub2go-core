package keyboards

import (
	"github.com/go-telegram/bot/models"
)

func ButtonBack() []models.InlineKeyboardButton {
	return []models.InlineKeyboardButton{{Text: "⬅️ Назад", CallbackData: "back"}}
}
