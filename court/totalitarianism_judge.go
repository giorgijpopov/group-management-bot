package court

import (
	"fmt"
	"time"

	"github.com/giorgijpopov/telebot"
)

func judgeTotalitarian(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error {
	if !materials.HasNudes {
		return nil
	}
	until := time.Now().Add(40 * time.Second)
	err := bot.Restrict(message.Chat, &telebot.ChatMember{
		Rights:          DicksRestrictedRights(),
		User:            message.Sender,
		Role:            "Admin",
		Title:           "Title",
		RestrictedUntil: until.Unix(),
	})
	if err != nil {
		return err
	}
	_, err = bot.Send(message.Chat, fmt.Sprintf("%s, you have been restricted until %v!", message.Sender.FirstName, until.Format(time.RFC822)), &telebot.SendOptions{
		ReplyTo: message,
	})
	return err
}
