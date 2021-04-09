package court

import (
	"fmt"
	"time"

	"github.com/giorgijpopov/telebot"
	"github.com/group-management-bot/poll"
)

func judgeDemocratically(bot *telebot.Bot, defendantMsg *telebot.Message, materials CaseMaterials) error {
	if !materials.HasNudes {
		return nil
	}

	question := fmt.Sprintf("It seems that %s has sent some nudes. What shoud we do with him?", defendantMsg.Sender.FirstName)
	_, err := poll.RunPoll(bot, defendantMsg, time.Minute, question,
		poll.NewDummyExecutor("Nothing"),
		poll.NewOnlyMessagesRestrictor(40*time.Second),
		poll.NewOnlyMessagesRestrictor(2*time.Minute),
		poll.NewOnlyMessagesRestrictor(time.Hour),
		poll.NewOnlyMessagesRestrictor(3*time.Hour),
	)
	return err
}
