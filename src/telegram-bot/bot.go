package bot

import (
	"log"

	"../api"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type telegramBot struct {
	handler api.MessageHandler
}

func New(handler api.MessageHandler) telegramBot {
	bot := telegramBot{}
	bot.handler = handler

	return bot
}

func (b *telegramBot) Process(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf(bot.Self.UserName)

	var channel tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	channel.Timeout = 60
	updates, err := bot.GetUpdatesChan(channel)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		UserName := update.Message.From.UserName
		ChatID := update.Message.Chat.ID
		Text := update.Message.Text
		log.Printf("[%s] %d %s", UserName, ChatID, Text)

		text := b.handler.Handle(Text)
		msg := tgbotapi.NewMessage(ChatID, text)

		bot.Send(msg)
	}
}
