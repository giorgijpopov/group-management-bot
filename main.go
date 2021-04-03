package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "github.com/group-management-bot/misc/telebot"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".

		Token:  os.Getenv("TELEBOT_SECRET"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Chat, fmt.Sprintf("Hello %s!", m.Sender.FirstName))
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Chat, fmt.Sprintf("%s, I don't understand you yet", m.Sender.FirstName))
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		b.Send(m.Chat, fmt.Sprintf("%s, be careful with pictures!!!", m.Sender.FirstName))
	})

	b.Handle(tb.OnChannelPost, func (m *tb.Message) {
		b.Send(m.Chat, fmt.Sprintf("%s, language!!!", m.Sender.FirstName))
	})

	b.Start()
}