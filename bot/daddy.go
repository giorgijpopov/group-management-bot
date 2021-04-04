package bot

import "github.com/giorgijpopov/telebot"

type daddy struct {
	id string
}

var _ telebot.Recipient = &daddy{}

func newDaddy(id string) *daddy {
	return &daddy{id: id}
}

func (d daddy) Recipient() string {
	return d.id
}
