package court

import telebot "gopkg.in/telebot.v3"

type Court interface {
	Judge(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error
}
