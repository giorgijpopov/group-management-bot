package court

import telebot "gopkg.in/telebot.v3"

type Regime int

const (
	Democracy       Regime = 0
	Totalitarianism Regime = 1
	Gogi            Regime = 2
)

type judge func(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error

var judgeByRegime = map[Regime]judge{
	Democracy:       judgeDemocratically,
	Totalitarianism: judgeTotalitarian,
	Gogi:            judgeGogi,
}
