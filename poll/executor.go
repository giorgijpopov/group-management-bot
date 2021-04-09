package poll

import (
	"time"

	"github.com/giorgijpopov/errorx"
	"github.com/giorgijpopov/telebot"
)

type OptionExecutor interface {
	Description() string
	Execute(bot *telebot.Bot, params ExecutorParams) error
}

type ExecutorParams struct {
	SourceMsg *telebot.Message
}

func RunPoll(bot *telebot.Bot, sourceMsg *telebot.Message, pollDuration time.Duration, question string, executors ...OptionExecutor) (*time.Timer, error) {
	if len(executors) == 0 {
		return nil, nil
	}

	poll := &telebot.Poll{
		Type:     telebot.PollRegular,
		Question: question,
	}

	for _, executor := range executors {
		poll.AddOptions(executor.Description())
	}

	pollMsg, err := bot.Send(sourceMsg.Chat, poll)
	if err != nil {
		return nil, err
	}

	params := ExecutorParams{
		SourceMsg: sourceMsg,
	}

	return time.AfterFunc(pollDuration, func() {
		handlePollResult(bot, pollMsg, params, executors...)
	}), nil
}

func handlePollResult(bot *telebot.Bot, pollMsg *telebot.Message, params ExecutorParams, executors ...OptionExecutor) {
	p, err := bot.StopPoll(pollMsg)
	if err != nil {
		panic(err)
	}

	best := executors[0]
	max := 0
	for i := range p.Options {
		if i >= len(executors) {
			panic(errorx.IllegalState.New("poll executor numbers (%d) less then poll options (%d)", len(executors), len(p.Options)))
		}

		if p.Options[i].VoterCount > max {
			best = executors[i]
		}
	}

	err = best.Execute(bot, params)
	if err != nil {
		panic(err)
	}
}
