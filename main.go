package main

import (
	"log"
	"rcselect-bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Конфиг загружен, токен получен:", cfg.TelegramToken != "")

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() && update.Message.Command() == "start" {
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("💻 Вопрос по работе ПО", "software"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("🔑 Проблемы с лицензией", "license"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("✏️ Задать свой вопрос", "other"),
				),
			)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "👋 Добро пожаловать в службу поддержки RCSelect!\n\nВыберите категорию обращения:")
			msg.ReplyMarkup = keyboard
			log.Printf("Сообщение от %s: %s", update.Message.From.UserName, update.Message.Text)
			if _, err := bot.Send(msg); err != nil {
				log.Println("ошибка отправки:", err)
			}
		}
	}
}
