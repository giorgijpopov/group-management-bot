package main

import (
	"log"
	"os"

	"github.com/group-management-bot/bot"
	"github.com/group-management-bot/court"
	"github.com/group-management-bot/nudespolice"
)

func main() {
	botManager, err := bot.NewBotManager(
		os.Getenv("TBOT_SECRET"),
		os.Getenv("TBOT_DADDY_ID"),
		nudespolice.NewPoliceman(),
		court.NewCourt(court.Totalitarianism),
	)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	botManager.SetupHandles()
	botManager.Start()
}
