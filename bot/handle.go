package bot

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	telebot "gopkg.in/telebot.v3"
	"github.com/group-management-bot/poll"
)

func promoteTo(bot *telebot.Bot, message *telebot.Message) error {
	user, err := extractSourceUser(bot, message)
	if err != nil || user == nil {
		return err
	}

	member, err := bot.ChatMemberOf(message.Chat, user)
	if err != nil {
		return err
	}
	if member.Role == telebot.Creator {
		_, err := bot.Send(message.Chat, "Can't promote the owner", &telebot.SendOptions{
			ReplyTo: message,
		})
		return err
	}

	// do not allow promote yourself if you are restricted
	if user.ID == message.Sender.ID && member.Role != telebot.Administrator {
		_, err := bot.Send(message.Chat, "You don't have admin rights!", &telebot.SendOptions{
			ReplyTo: message,
		})
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

	member, err := bot.ChatMemberOf(message.Chat, user)
	if err != nil {
		return err
	}
	if member.Role == telebot.Creator {
		_, err := bot.Send(message.Chat, "Can't ban the owner", &telebot.SendOptions{
			ReplyTo: message,
		})
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

func rollDice(bot *telebot.Bot, message *telebot.Message) error {
	// Check if command is sent by specific usernames
	username := message.Sender.Username
	if username == "dnzonzor" || username == "q1ruwa" {
		_, err := bot.Send(message.Chat, "дд, анох, пошли нахуй!", &telebot.SendOptions{
			ReplyTo: message,
		})
		return err
	}

	// Parse arguments: /roll [min] [max]
	// Default: 1 to 6
	min := 1
	max := 6

	payload := strings.TrimSpace(message.Payload)
	if payload != "" {
		args := strings.Fields(payload)

		if len(args) == 1 {
			// /roll 10 -> rolls 1 to 10
			if val, err := strconv.Atoi(args[0]); err == nil && val > 0 {
				max = val
			} else {
				_, err := bot.Send(message.Chat, "❌ Invalid number. Use: /roll [max] or /roll [min] [max]", &telebot.SendOptions{
					ReplyTo: message,
				})
				return err
			}
		} else if len(args) >= 2 {
			// /roll 1 10 -> rolls 1 to 10
			minVal, err1 := strconv.Atoi(args[0])
			maxVal, err2 := strconv.Atoi(args[1])

			if err1 != nil || err2 != nil || minVal >= maxVal {
				_, err := bot.Send(message.Chat, "❌ Invalid range. Use: /roll [min] [max] where min < max", &telebot.SendOptions{
					ReplyTo: message,
				})
				return err
			}

			min = minVal
			max = maxVal
		}
	}

	// Validate range
	if max-min > 1000000 {
		_, err := bot.Send(message.Chat, "❌ Range too large! Maximum range is 1,000,000", &telebot.SendOptions{
			ReplyTo: message,
		})
		return err
	}

	// Roll the dice
	result := rand.Intn(max-min+1) + min

	var response string
	if min == 1 && max == 6 {
		response = fmt.Sprintf("🎲 %s rolled: %d", message.Sender.FirstName, result)
	} else {
		response = fmt.Sprintf("🎲 %s rolled (%d-%d): %d", message.Sender.FirstName, min, max, result)
	}

	_, err := bot.Send(message.Chat, response, &telebot.SendOptions{
		ReplyTo: message,
	})
	return err
}
