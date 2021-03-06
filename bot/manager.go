package bot

import (
	"image"
	"log"
	"time"

	"github.com/giorgijpopov/telebot"
	"github.com/group-management-bot/court"
	"github.com/group-management-bot/nudespolice"
)

const (
	pollerTimeout = 10 * time.Second
)

type manager struct {
	bot   *telebot.Bot
	daddy *daddy

	nudePoliceman nudespolice.Policeman
	court         court.Court
}

var _ Manager = &manager{}

func NewBotManager(
	token string,
	daddyID string,
	nudePoliceman nudespolice.Policeman,
	court court.Court,
) (*manager, error) {

	m := &manager{
		nudePoliceman: nudePoliceman,
		court:         court,
		daddy:         newDaddy(daddyID),
	}

	poller := &telebot.LongPoller{Timeout: pollerTimeout}
	privateMessagesForbiddenMiddleware := telebot.NewMiddlewarePoller(poller, func(upd *telebot.Update) bool {
		if upd.Message == nil {
			return true
		}

		if upd.Message.Chat.Type == telebot.ChatPrivate {
			return false
		}

		return true
	})
	b, err := telebot.NewBot(telebot.Settings{
		Token:    token,
		Poller:   privateMessagesForbiddenMiddleware,
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
	m.bot.Handle(telebot.OnPhoto, m.defaultHandler)
	m.bot.Handle(telebot.OnDocument, m.defaultHandler)
	m.bot.Handle("/promoteTo", m.promoteTo)
	m.bot.Handle("/banFor", m.banFor)
}

func (m *manager) promoteTo(message *telebot.Message) {
	err := promoteTo(m.bot, message)
	m.HandleError(err)
}

func (m *manager) banFor(message *telebot.Message) {
	err := banFor(m.bot, message)
	m.HandleError(err)
}

func (m *manager) defaultHandler(message *telebot.Message) {
	caseMaterials, err := m.gatherCaseMaterials(message)
	if !m.HandleError(err) {
		return
	}
	m.HandleError(m.court.Judge(m.bot, message, caseMaterials))
}

func (m *manager) findImageInMessage(message *telebot.Message) (image.Image, error) {
	var file telebot.File
	switch {
	case message.Photo != nil:
		file = message.Photo.File
	case message.Document != nil:
		file = message.Document.File
	default:
		return nil, nil
	}

	reader, err := m.bot.GetFile(&file)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (m *manager) gatherCaseMaterials(message *telebot.Message) (court.CaseMaterials, error) {
	res := court.CaseMaterials{}

	img, err := m.findImageInMessage(message)
	if err != nil {
		return court.CaseMaterials{}, err
	}

	if img != nil {
		res.HasNudes, err = m.nudePoliceman.CheckNudesInImage(img)
		if err != nil {
			return court.CaseMaterials{}, err
		}
	}

	return res, nil
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
	_, err := m.bot.Send(m.daddy, complaint)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
