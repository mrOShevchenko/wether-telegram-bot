package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

const (
	commandStartMessage = "OK, let's see what holiday today at this countries:"
	greetingMessage     = " Hello I am Name Bot. Press /start to see country's flags"
)

func (tg *TelegramBot) PrintMessage(message *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	if text == commandStartMessage {
		holidays, err := tg.HolidayRequest()
		if err != nil {
			return errors.Wrapf(err, "error in HolidayRequest")
		}

		msg.ReplyMarkup = tg.CountrySelection(holidays)
	}
	if err := tg.SendMessage(msg); err != nil {
		return errors.Wrap(err, "error with sending message")
	}
	return nil
}

func (tg *TelegramBot) SendMessage(msg tgbotapi.MessageConfig) error {
	if _, err := tg.bot.Send(msg); err != nil {
		return errors.Wrap(err, "send message failed")
	}
	return nil
}
