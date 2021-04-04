package poll

import (
	"github.com/giorgijpopov/telebot"
)

type DummyExecutor struct {
	descr string
}

var _ OptionExecutor = DummyExecutor{}

func NewDummyExecutor(descr string) DummyExecutor {
	return DummyExecutor{descr: descr}
}

func (p DummyExecutor) Description() string {
	return p.descr
}

func (p DummyExecutor) Execute(*telebot.Bot, ExecutorParams) error {
	return nil
}
