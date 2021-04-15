package poll

import (
	"fmt"
	"time"

	"github.com/giorgijpopov/telebot"
)

type OnlyMessagesRestrictor struct {
	duration time.Duration
}

var _ OptionExecutor = OnlyMessagesRestrictor{}

func NewOnlyMessagesRestrictor(duration time.Duration) OnlyMessagesRestrictor {
	return OnlyMessagesRestrictor{duration: duration}
}

func (p OnlyMessagesRestrictor) Description() string {
	return fmt.Sprintf("Restrict sending anything but messages for %v", p.duration)
}

func (p OnlyMessagesRestrictor) Execute(bot *telebot.Bot, params ExecutorParams) error {
	until := time.Now().Add(p.duration)
	err := bot.Restrict(params.Chat, &telebot.ChatMember{
		Rights:          OnlyMessagesRights(),
		User:            params.User,
		RestrictedUntil: until.Unix(),
	})
	if err != nil {
		return err
	}

	_, err = bot.Send(params.Chat, fmt.Sprintf("%s, you have been restricted until %v!", params.User.FirstName, until.Format(time.RFC822)), &telebot.SendOptions{
		ReplyToID: params.SourceMessageID,
	})
	return err
}

func OnlyMessagesRights() telebot.Rights {
	return telebot.Rights{
		CanSendMessages: true,
	}
}
