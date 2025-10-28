package court

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
)

func judgeGogi(bot *telebot.Bot, defendantMsg *telebot.Message, materials CaseMaterials) error {
	message := fmt.Sprintf("It seems that %s has sent some nudes. ⚠️ ГОГИ, ПОШЕЛ НАХУЙ! Выключи эту фичу!", defendantMsg.Sender.FirstName)

	_, err := bot.Send(defendantMsg.Chat, message, &telebot.SendOptions{
		ReplyTo: defendantMsg,
	})
	return err
}
