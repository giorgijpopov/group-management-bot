package bot

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/giorgijpopov/telebot"
	"github.com/koyachi/go-nude"
)

const (
	pollerTimeout = 10 * time.Second
)

type manager struct {
	bot *telebot.Bot
}

var _ Manager = &manager{}

func NewBotManager(tocken string) (*manager, error) {
	m := &manager{}
	b, err := telebot.NewBot(telebot.Settings{
		Token:    tocken,
		Poller:   &telebot.LongPoller{Timeout: pollerTimeout},
		Reporter: m.reportError,
	})
	if err != nil {
		return nil, err
	}
	m.bot = b
	return m, nil
}

func (m *manager) Start() {
	m.bot.Start()
}

func (m *manager) SetupHandles() {
	m.bot.Handle(telebot.OnPhoto, func(message *telebot.Message) {
		reader, err := m.bot.GetFile(&message.Photo.File)
		if !m.HandleError(err) {
			return
		}

		img, _, err := image.Decode(reader)
		if !m.HandleError(err) {
			return
		}

		hasNudes, err := nude.IsImageNude(img)
		if !m.HandleError(err) {
			return
		}

		if !hasNudes {
			return
		}

		until := time.Now().Add(40 * time.Second)
		err = m.bot.Restrict(message.Chat, &telebot.ChatMember{
			Rights:          DicksRestrictedRights(),
			User:            message.Sender,
			Role:            "Admin",
			Title:           "Title",
			RestrictedUntil: until.Unix(),
		})
		if !m.HandleError(err) {
			return
		}
		_, err = m.bot.Send(message.Chat, fmt.Sprintf("%s, you have been restricted until %v!", message.Sender.FirstName, until.Format(time.RFC822)), &telebot.SendOptions{
			ReplyTo: message,
		})
		m.HandleError(err)
	})
}

func (m *manager) HandleError(err error) bool {
	if err != nil {
		m.reportError(err)
		return false
	}
	return true
}

func (m *manager) reportError(err error) {
	m.complainToDaddy(err.Error())
}

func (m *manager) complainToDaddy(complaint string) {
	_, err := m.bot.Send(m.bot.Me, complaint)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}

// DicksRestrictedRights allow to send only messages.
func DicksRestrictedRights() telebot.Rights {
	return telebot.Rights{
		CanSendMessages: true,
	}
}
