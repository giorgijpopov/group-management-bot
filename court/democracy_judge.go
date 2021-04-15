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

	pollExecutorParams := poll.ExecutorParams{
		Chat:            defendantMsg.Chat,
		User:            defendantMsg.Sender,
		SourceMessageID: defendantMsg.ID,
		PollDuration:    time.Minute,
		Question:        question,
	}
	err := poll.RunPoll(bot, pollExecutorParams,
		poll.NewDummyExecutor("Nothing"),
		poll.NewOnlyMessagesRestrictor(5*time.Minute),
		poll.NewOnlyMessagesRestrictor(20*time.Minute),
		poll.NewOnlyMessagesRestrictor(time.Hour),
		poll.NewOnlyMessagesRestrictor(3*time.Hour),
	)
	return err
}
