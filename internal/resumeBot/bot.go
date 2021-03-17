package resumeBot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var change map[int64][]tgbotapi.Message = make(map[int64][]tgbotapi.Message)

var bot *tgbotapi.BotAPI

func Start(token string) error {
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	newUC := tgbotapi.NewUpdate(0)
	newUC.Timeout = 40

	uChan, err := bot.GetUpdatesChan(newUC)
	if err != nil {
		return err
	}
	for update := range uChan {
		if update.Message != nil {
			handleMessage(update)
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(update)
		}
	}
	return nil
}
