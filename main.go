package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var menuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Books"),
		tgbotapi.NewKeyboardButton("Games"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Videos"),
		tgbotapi.NewKeyboardButton("Jokes"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("GOLANGTZBOT"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates.
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages for now. Will add filtering etc later.
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "So far I can only make Jokes. Use `/joke`"
		case "hi":
			msg.Text = "Hello there! :)"
		case "status":
			msg.Text = "I'm incomplete :("
		case "joke":
			msg.Text = Joke()
		case "menu":
			msg.ReplyMarkup = menuKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
