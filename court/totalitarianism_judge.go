package court

import (
	"github.com/giorgijpopov/telebot"
	"github.com/group-management-bot/poll"
)

func judgeTotalitarian(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error {
	if !materials.HasNudes {
		return nil
	}

	return poll.NewOnlyMessagesRestrictor(0).Execute(bot, poll.ExecutorParams{SourceMsg: message})
}
