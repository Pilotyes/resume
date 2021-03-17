package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func WelcomeMessage(chatID int64) tgbotapi.Chattable {
	welcomeMessage := "Привет!\nЭто телеграм-резюме-бот.\nЗдесь представлена основная информация обо мне, моём опыте и ещё несколько интересных фактов."
	buttonAboutMe := tgbotapi.NewInlineKeyboardButtonData("Давай познакомимся поближе!", "/aboutme")
	buttonsRowAboutMe := tgbotapi.NewInlineKeyboardRow(buttonAboutMe)
	messageMarkup := tgbotapi.NewInlineKeyboardMarkup(buttonsRowAboutMe)
	message := tgbotapi.NewMessage(chatID, welcomeMessage)
	message.ReplyMarkup = messageMarkup
	return message
}
