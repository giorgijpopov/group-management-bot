package court

import "github.com/giorgijpopov/telebot"

type CaseMaterials struct {
	HasNudes bool
}

// DicksRestrictedRights allow to send only messages.
func DicksRestrictedRights() telebot.Rights {
	return telebot.Rights{
		CanSendMessages: true,
	}
}
