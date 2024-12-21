package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Проверяем, является ли сообщение из нужного канала
		// test new commit
		if update.Message.Chat.Type == tgbotapi.ChatTypeChannel && update.Message.Chat.Title == "МойКанал" {
			// Извлекаем необходимую информацию из сообщения
			text := update.Message.Text

			// Проверяем, содержит ли сообщение ключевое слово "заявка"
			if strings.Contains(text, "заявка") {
				// Создаем заявку (здесь нужно реализовать логику отправки заявки)
				fmt.Printf("Создана заявка: %s\n", text)
			}
		}
	}
}
