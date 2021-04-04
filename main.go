package main

import (
	"log"
	"os"

	"github.com/group-management-bot/bot"
)

func main() {
	botManager, err := bot.NewBotManager(os.Getenv("TELEBOT_SECRET"))
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	botManager.SetupHandles()
	botManager.Start()
}
