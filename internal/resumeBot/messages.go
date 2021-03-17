package resumeBot

import (
	"resume/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleMessage(update tgbotapi.Update) {
	switch update.Message.Text {

	case "/start":
		message := model.WelcomeMessage(update.Message.Chat.ID)
		sendMessage(message, update.Message.Chat.ID)
	}
}
