package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"time"

	tb "github.com/group-management-bot/misc/telebot"
	nude "github.com/koyachi/go-nude"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEBOT_SECRET"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		reader, err := b.GetFile(&m.Photo.File)
		if err != nil {
			logError(err)
			return
		}
		img, _, err := image.Decode(reader)
		if err != nil {
			logError(err)
			return
		}
		hasNudes, err := nude.IsImageNude(img)
		if err != nil {
			logError(err)
			return
		}
		if !hasNudes {
			return
		}
		until := time.Now().Add(time.Minute)
		err = b.Restrict(m.Chat, &tb.ChatMember{
			Rights:          DicksRestrictedRights(),
			User:            m.Sender,
			Role:            "Admin",
			Title:           "Title",
			RestrictedUntil: until.Unix(),
		})
		if err != nil {
			logError(err)
			return
		}
		_, err = b.Send(m.Chat, fmt.Sprintf("%s, you have been restricted until %v!", m.Sender.FirstName, until.Format(time.RFC822)), &tb.SendOptions{
			ReplyTo: m,
		})
		logError(err)
	})

	b.Start()
}

func logError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// DicksRestrictedRights allow to send only messages.
func DicksRestrictedRights() tb.Rights {
	return tb.Rights{
		CanSendMessages: true,
	}
}
