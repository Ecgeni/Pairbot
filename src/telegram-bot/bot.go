package bot

import (
	"log"

	"../api"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Exchange struct {
	PairCharge []api.GetPairData
}

func (e *Exchange) Process(token string) {
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

		text := e.selectMessage(Text)
		msg := tgbotapi.NewMessage(ChatID, text)

		bot.Send(msg)
	}
}

func (e *Exchange) selectMessage(text string) string {
	var result string
	switch text {
	case "/help":
		result = "Available commands: \n /show - Show all available trade pairs."
		break
	case "/show":
		for _, item := range e.PairCharge {
			result += item.Content() + "\n"
		}
		break
	default:
		result = "Not available command."
	}

	return result
}
