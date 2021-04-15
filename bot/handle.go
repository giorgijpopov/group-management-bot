package bot

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/giorgijpopov/telebot"
	"github.com/group-management-bot/poll"
)

func promoteTo(bot *telebot.Bot, message *telebot.Message) error {
	user, err := extractSourceUser(bot, message)
	if err != nil || user == nil {
		return err
	}

	title := message.Payload
	if title == "" {
		_, err := bot.Send(message.Chat, "Set title after the command")
		return err
	}

	question := fmt.Sprintf("%s suggests to promote %s and set him title %s", message.Sender.FirstName, user.FirstName, title)

	pollExecutorParams := poll.ExecutorParams{
		Chat:         message.Chat,
		User:         user,
		PollDuration: time.Minute,
		Question:     question,
	}
	return poll.RunPoll(bot, pollExecutorParams,
		poll.NewDummyExecutor("Nope"),
		poll.NewPromoter(title),
	)
}

func banFor(bot *telebot.Bot, message *telebot.Message) error {
	user, err := extractSourceUser(bot, message)
	if err != nil || user == nil {
		return err
	}

	explanation := "Use number + m (minutes) or h (hours). For example 20m or 1h"

	input := message.Payload
	var durationUnit time.Duration
	switch {
	case strings.Contains(input, "m"):
		durationUnit = time.Minute
	case strings.Contains(input, "h"):
		durationUnit = time.Hour
	default:
		_, err := bot.Send(message.Chat, explanation)
		return err
	}

	num, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		_, err := bot.Send(message.Chat, explanation)
		return err
	}

	question := fmt.Sprintf("%s suggests to restrict %s for %s", message.Sender.FirstName, user.FirstName, input)

	pollExecutorParams := poll.ExecutorParams{
		Chat:         message.Chat,
		User:         user,
		PollDuration: time.Minute,
		Question:     question,
	}
	return poll.RunPoll(bot, pollExecutorParams,
		poll.NewDummyExecutor("Nope"),
		poll.NewOnlyMessagesRestrictor(time.Duration(num)*durationUnit),
	)
}

func extractSourceUser(bot *telebot.Bot, message *telebot.Message) (*telebot.User, error) {
	if !message.IsReply() {
		_, err := bot.Send(message.Chat, "Message must be reply")
		return nil, err
	}
	return message.ReplyTo.Sender, nil
}
