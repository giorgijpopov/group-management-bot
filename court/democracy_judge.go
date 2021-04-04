package court

import (
	"fmt"
	"time"

	"github.com/giorgijpopov/telebot"
)

func judgeDemocratically(bot *telebot.Bot, defendantMsg *telebot.Message, materials CaseMaterials) error {
	if !materials.HasNudes {
		return nil
	}

	poll := &telebot.Poll{
		Type:     telebot.PollRegular,
		Question: fmt.Sprintf("It seems that %s has sent some nudes. What shoud we do with him?", defendantMsg.Sender.FirstName),
	}
	poll.AddOptions("Nothing", "Restrict sending anything but messages for 40 sec")

	pollMsg, err := bot.Send(defendantMsg.Chat, poll)
	if err != nil {
		return err
	}

	go handlePollResult(bot, pollMsg, defendantMsg)
	return nil
}

func handlePollResult(bot *telebot.Bot, pollMsg, defendantMsg *telebot.Message) {
	time.Sleep(time.Minute)

	p, err := bot.StopPoll(pollMsg)
	if err != nil {
		panic(err)
	}

	if p.Options[0].VoterCount > p.Options[1].VoterCount {
		return
	}

	until := time.Now().Add(40 * time.Second)
	err = bot.Restrict(defendantMsg.Chat, &telebot.ChatMember{
		Rights:          DicksRestrictedRights(),
		User:            defendantMsg.Sender,
		Role:            "Admin",
		Title:           "Title",
		RestrictedUntil: until.Unix(),
	})
	if err != nil {
		panic(err)
	}

	_, err = bot.Send(defendantMsg.Chat, fmt.Sprintf("%s, you have been restricted until %v!", defendantMsg.Sender.FirstName, until.Format(time.RFC822)), &telebot.SendOptions{
		ReplyTo: defendantMsg,
	})
}
