package keyboards

import (
	"github.com/go-telegram/bot/models"
)

func AdminPanelKeyboard() *models.InlineKeyboardMarkup {
	var keyboard [][]models.InlineKeyboardButton

	buttonBack := ButtonBack()

	keyboard = append(keyboard,
		[]models.InlineKeyboardButton{
			{Text: "🟢 Добавить пользователя", CallbackData: "add_user"},
			{Text: "🔴 Удалить пользователя", CallbackData: "del_user"}},

		[]models.InlineKeyboardButton{
			{Text: "🔔 Добавить подписку", CallbackData: "add_sub"},
			{Text: "🔕 Удалить подписку", CallbackData: "del_sub"}},

		buttonBack,
	)

	return &models.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}
