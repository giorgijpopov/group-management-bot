package court

import "github.com/giorgijpopov/telebot"

type Court interface {
	Judge(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error
}
