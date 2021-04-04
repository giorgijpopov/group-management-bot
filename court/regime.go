package court

import "github.com/giorgijpopov/telebot"

type Regime int

const (
	Democracy       Regime = 0
	Totalitarianism Regime = 1
)

type judge func(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error

var judgeByRegime = map[Regime]judge{
	Democracy:       judgeDemocratically,
	Totalitarianism: judgeTotalitarian,
}
