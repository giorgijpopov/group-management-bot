package court

import (
	"github.com/group-management-bot/poll"
	telebot "gopkg.in/telebot.v3"
)

func judgeTotalitarian(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error {
	if !materials.HasNudes {
		return nil
	}

	return poll.NewOnlyMessagesRestrictor(0).Execute(bot, poll.ExecutorParams{
		Chat:            message.Chat,
		User:            message.Sender,
		SourceMessageID: message.ID,
	})
}
