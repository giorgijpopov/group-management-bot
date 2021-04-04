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
	until := time.Now().Add(40 * time.Second)
	err := bot.Restrict(params.SourceMsg.Chat, &telebot.ChatMember{
		Rights:          OnlyMessagesRights(),
		User:            params.SourceMsg.Sender,
		RestrictedUntil: until.Unix(),
	})
	if err != nil {
		return err
	}

	_, err = bot.Send(params.SourceMsg.Chat, fmt.Sprintf("%s, you have been restricted until %v!", params.SourceMsg.Sender.FirstName, until.Format(time.RFC822)), &telebot.SendOptions{
		ReplyTo: params.SourceMsg,
	})
	return err
}

func OnlyMessagesRights() telebot.Rights {
	return telebot.Rights{
		CanSendMessages: true,
	}
}
