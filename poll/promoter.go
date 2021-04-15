package poll

import (
	"fmt"

	"github.com/giorgijpopov/telebot"
)

type Promoter struct {
	title string
}

var _ OptionExecutor = Promoter{}

func NewPromoter(title string) Promoter {
	return Promoter{title: title}
}

func (p Promoter) Description() string {
	return fmt.Sprintf("Promote")
}

func (p Promoter) Execute(bot *telebot.Bot, params ExecutorParams) error {
	err := bot.Promote(params.Chat, &telebot.ChatMember{
		Rights: AdminRights(),
		User:   params.User,
	})
	if err != nil {
		return err
	}

	return bot.SetAdminTitle(params.Chat, params.User, p.title)
}

func AdminRights() telebot.Rights {
	rights := telebot.AdminRights()
	rights.CanPostMessages = false
	rights.CanEditMessages = false
	return rights
}
