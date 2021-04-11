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
	Chat            *telebot.Chat
	User            *telebot.User
	SourceMessageID int
	PollDuration    time.Duration
	Question        string
}

func RunPoll(bot *telebot.Bot, params ExecutorParams, executors ...OptionExecutor) error {
	handleFunc, err := createPollAndGetCallback(bot, params, executors...)
	if err != nil {
		return err
	}

	time.Sleep(params.PollDuration)
	return handleFunc()
}

func RunPollAsync(bot *telebot.Bot, params ExecutorParams, executors ...OptionExecutor) (*time.Timer, error) {
	handleFunc, err := createPollAndGetCallback(bot, params, executors...)
	if err != nil {
		return nil, err
	}

	return time.AfterFunc(params.PollDuration, func() {
		err := handleFunc()
		if err != nil {
			panic(err)
		}
	}), nil
}

func createPollAndGetCallback(bot *telebot.Bot, params ExecutorParams, executors ...OptionExecutor) (func() error, error) {
	if len(executors) == 0 {
		return nil, nil
	}

	poll := &telebot.Poll{
		Type:     telebot.PollRegular,
		Question: params.Question,
	}

	for _, executor := range executors {
		poll.AddOptions(executor.Description())
	}

	pollMsg, err := bot.Send(params.Chat, poll)
	if err != nil {
		return nil, err
	}

	return func() error {
		return handlePollResult(bot, pollMsg, params, executors...)
	}, nil
}

func handlePollResult(bot *telebot.Bot, pollMsg *telebot.Message, params ExecutorParams, executors ...OptionExecutor) error {
	p, err := bot.StopPoll(pollMsg)
	if err != nil {
		return err
	}

	best := executors[0]
	max := 0
	for i := range p.Options {
		if i >= len(executors) {
			return errorx.IllegalState.New("poll executor numbers (%d) less then poll options (%d)", len(executors), len(p.Options))
		}

		if p.Options[i].VoterCount > max {
			max = p.Options[i].VoterCount
			best = executors[i]
		}
	}

	err = best.Execute(bot, params)
	if err != nil {
		return err
	}
	return nil
}
