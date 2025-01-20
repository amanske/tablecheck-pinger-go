package telebot

import (
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/telebot.v4"
	tele "gopkg.in/telebot.v4"
)

func Initialize() *telebot.Bot {

	// Retrieve the API token using os.Getenv
	apiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if apiToken == "" {
		log.Fatal("API token not found in environment variables. Please make sure you have set TELEGRAM_BOT_API_TOKEN.")
	}

	// Initialize the bot with the retrieved API token
	bot, err := tele.NewBot(
		telebot.Settings{
			Token:  apiToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func Recipient(channelID string) (telebot.Recipient, error) {
	channel, err := strconv.Atoi(channelID)
	if err != nil {
		return nil, err
	}
	return telebot.ChatID(channel), nil
}
