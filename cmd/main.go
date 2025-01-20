package main

import (
	"fmt"
	"log"
	"os"
	"tablecheck-pinger-go/internal/query"
	bot "tablecheck-pinger-go/telebot"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: tablecheck-pinger-go <url> <channel_username>")
	}

	url := os.Args[1]
	channelID := os.Args[2]

	b := bot.Initialize()
	channel, err := bot.Recipient(channelID)
	if err != nil {
		log.Fatal(err)
	}

	found := query.CheckAvailability(url)

	if found {
		_, err := b.Send(channel, fmt.Sprintf("Found an available slot for %s\n", url))
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		fmt.Println("Available slots found!")
	} else {
		fmt.Println("No available slots found")
	}
}
