package court

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)

type court struct {
	regime Regime
}

var _ Court = &court{}

func NewCourt(regime Regime) *court {
	return &court{regime: regime}
}

func (c *court) Judge(bot *telebot.Bot, message *telebot.Message, materials CaseMaterials) error {
	if !materials.HasNudes {
		return nil
	}

	judge, found := judgeByRegime[c.regime]
	if !found {
		return fmt.Errorf("not existent regime %d", c.regime)
	}

	return judge(bot, message, materials)
}
