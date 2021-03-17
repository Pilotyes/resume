package resumeBot

import (
	"resume/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func deleteLastMessage(chatID int64) error {
	length := len(change[chatID])
	if length < 2 {
		return nil
	}
	bot.DeleteMessage(tgbotapi.DeleteMessageConfig{ChatID: chatID, MessageID: change[chatID][length-1].MessageID})
	//newMessage := tgbotapi.NewMessage(chatID, change[chatID][length-2].Text)
	//newMessage.ReplyMarkup = change[chatID][length-2].ReplyMarkup
	//bot.Send(newMessage)
	change[chatID] = change[chatID][:length-1]
	return nil
}

func sendMessage(message tgbotapi.Chattable, chatID int64) error {
	sampleMess, err := bot.Send(message)
	change[chatID] = append(change[chatID], sampleMess)
	return err
}

func handleCallbackQuery(update tgbotapi.Update) {

	chatID := update.CallbackQuery.Message.Chat.ID

	switch update.CallbackQuery.Data {
	case "/aboutme":
		deleteLastMessage(chatID)

		aboutMe := tgbotapi.NewPhotoUpload(chatID, config.PATH_TO_PHOTOS_DIR+"1.jpg")
		aboutMe.Caption = "Это я - Новосёлова Алёна. Мне 25 лет.\nМой родной город Ульяновск, но последние 2 года живу и работаю в Санкт-Петербурге.\n\nА сейчас я хочу рассказать..."

		buttonEducation := tgbotapi.NewInlineKeyboardButtonData("Об образовании", "/education")
		buttonExperience := tgbotapi.NewInlineKeyboardButtonData("Об опыте работы", "/experience")
		buttonHobbies := tgbotapi.NewInlineKeyboardButtonData("О хобби", "/hobbies")

		buttonsRowEducation := tgbotapi.NewInlineKeyboardRow(buttonEducation)
		buttonsRowExperience := tgbotapi.NewInlineKeyboardRow(buttonExperience)
		buttonsRowHobbies := tgbotapi.NewInlineKeyboardRow(buttonHobbies)

		messageMarkup := tgbotapi.NewInlineKeyboardMarkup(buttonsRowEducation, buttonsRowExperience, buttonsRowHobbies)
		aboutMe.ReplyMarkup = messageMarkup

		sendMessage(aboutMe, chatID)

	case "/education":
		deleteLastMessage(chatID)

		buttonMagister := tgbotapi.NewInlineKeyboardButtonData("Магистратуры", "/magister")
		buttonBachelor := tgbotapi.NewInlineKeyboardButtonData("Бакалавриата", "/bachelor")
		buttonOtherEducation := tgbotapi.NewInlineKeyboardButtonData("Переподготовки", "/other")
		buttonBack := tgbotapi.NewInlineKeyboardButtonData("Назад", "/back")

		buttonsRowMagister := tgbotapi.NewInlineKeyboardRow(buttonMagister)
		buttonsRowBachelor := tgbotapi.NewInlineKeyboardRow(buttonBachelor)
		buttonsRowOtherEducation := tgbotapi.NewInlineKeyboardRow(buttonOtherEducation)
		buttonsRowBack := tgbotapi.NewInlineKeyboardRow(buttonBack)

		messageMarkup := tgbotapi.NewInlineKeyboardMarkup(buttonsRowMagister, buttonsRowBachelor, buttonsRowOtherEducation, buttonsRowBack)
		informationAbout := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Образование - важная и довольно большая часть моей жизни, поэтому предлагаю начать с...")
		informationAbout.ReplyMarkup = messageMarkup

		sendMessage(informationAbout, chatID)

	case "/back":
		deleteLastMessage(chatID)

	case "/magister":
		deleteLastMessage(chatID)

		aboutMagister := tgbotapi.NewPhotoUpload(chatID, config.PATH_TO_PHOTOS_DIR+"2.jpg")
		aboutMagister.Caption = "Диплом Магистра - 2019 г.\nВУЗ: Ульяновский Государственный Университет.\nФакультет экономики. Направление: экономическая безопасность организации."

		buttonBack := tgbotapi.NewInlineKeyboardButtonData("Назад", "/back")
		buttonsRowBack := tgbotapi.NewInlineKeyboardRow(buttonBack)
		messageMarkup := tgbotapi.NewInlineKeyboardMarkup(buttonsRowBack)
		aboutMagister.ReplyMarkup = messageMarkup

		sendMessage(aboutMagister, chatID)

	case "/bachelor":
		deleteLastMessage(chatID)

		aboutBachelor := tgbotapi.NewPhotoUpload(chatID, config.PATH_TO_PHOTOS_DIR+"2.jpg")
		aboutBachelor.Caption = "Диплом Бакалавра - 2017 г.\nВУЗ: Ульяновский Институт Гражданской Авиации имени Б.П. Бугаева.\nФакультет подготовки авиационных специалистов. Направление: авиатопливное обеспечение воздушных перевозок и авиационных работ.\nСтаршина группы."

		buttonBack := tgbotapi.NewInlineKeyboardButtonData("Назад", "/back")
		buttonsRowBack := tgbotapi.NewInlineKeyboardRow(buttonBack)
		messageMarkup := tgbotapi.NewInlineKeyboardMarkup(buttonsRowBack)
		aboutBachelor.ReplyMarkup = messageMarkup

		sendMessage(aboutBachelor, chatID)

	case "/other":
		deleteLastMessage(chatID)

		aboutOther := tgbotapi.NewPhotoUpload(chatID, config.PATH_TO_PHOTOS_DIR+"2.jpg")
		aboutOther.Caption = "В процессе - 2021 г.\nКурсы: GeekBrains (Mail.ru)\nФакультет Golang разработки.\n\nДиплом переподготовки - 2019 г.\nВУЗ: БГТУ имени В.Г.Шухова.\nНаправление: техносферная безопасность.\n\nДиплом переподготовки - 2017 г.\nВУЗ: УлГПУ имени И.Н.Ульяноваю.\nНаправление: дополнительное образование детей."

		buttonBack := tgbotapi.NewInlineKeyboardButtonData("Назад", "/back")
		buttonsRowBack := tgbotapi.NewInlineKeyboardRow(buttonBack)
		messageMarkup := tgbotapi.NewInlineKeyboardMarkup(buttonsRowBack)
		aboutOther.ReplyMarkup = messageMarkup

		sendMessage(aboutOther, chatID)

	case "/experience":
		deleteLastMessage(chatID)

		aboutExperience := tgbotapi.NewPhotoUpload(chatID, config.PATH_TO_PHOTOS_DIR+"2.jpg")
		aboutExperience.Caption = "Компания Mars\nДолжность: специалист по охране труда и безопасности.\nС июля 2018 г. - по н.в.\n\nПримеры активностей:\n1. Лидирование и реализация процесса перехода на электронный документооборот.\n2. Реализация проекта по модернизации внутренней коммуникации отдела.\n3. Лидирование и реализация проекта по повышению Safety culture.\n4. Реализация проекта по созданию кросс-функциональных команд по выявлению и устранению рисков.\n5. Лидирование проекта по модернизации стандарта оценки рисков и работы с подрядными организациями.\n6. Лидирование проекта в области Health & Wellbeing.\n\n\nКомпания: Sollers\nДолжность: специалист по закупкам.\nС августа 2017 г. - по июнь 2018 г.\n\nПримеры активностей:\n1. Создание, ведение и обновление базы данных по поставщикам.\n2. Подготовка к согласованию и заключению договорных документов.\n3. Организация выбора поставщиков.\n4. Проведение оценки суммы затрат.\n5. Мониторинг рынка комплектующих изделий, поставщиков"

		sendMessage(aboutExperience, chatID)

	case "/hobbies":
		deleteLastMessage(chatID)

		aboutHobbies := tgbotapi.NewPhotoUpload(chatID, config.PATH_TO_PHOTOS_DIR+"2.jpg")
		aboutHobbies.Caption = "Моё самое любимое хобби - учиться.\nЯ изучаю английский (на данный момент мой уровень Intermediate), читаю развивающую литературу (Soft Skills, экономические и политические обзоры и т.д.), развиваю навыки в области программирования (Linux; Golang; Git; Docker; JavaScript; MySQL; HTML&CSS).\nВ свободное время играю в интеллектуальные игры, смотрю фильмы, путешествую."

		sendMessage(aboutHobbies, chatID)
	}
}
