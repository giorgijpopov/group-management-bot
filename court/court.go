package court

import (
	"github.com/giorgijpopov/errorx"
	"github.com/giorgijpopov/telebot"
)

type court struct {
	regime Regime
}

var _ Court = &court{}

func NewCourt(regime Regime) *court {
	return &court{regime: regime}
}

func (c *court) Judge(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error {
	judge, found := judgeByRegime[c.regime]
	if !found {
		return errorx.IllegalState.New("not existent regime %s", c.regime)
	}

	return judge(bot, message, materials)
}
