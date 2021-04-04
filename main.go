package main

import (
	"log"
	"os"

	"github.com/group-management-bot/bot"
	"github.com/group-management-bot/nudespolice"
)

func main() {
	botManager, err := bot.NewBotManager(os.Getenv("TELEBOT_SECRET"), nudespolice.NewPoliceman())
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	botManager.SetupHandles()
	botManager.Start()
}
