package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

const (
	commandStartMessage = "OK, let's see what holiday today at this countries:"
	greetingMessage     = " Hello I am Holiday Bot. Press /start to see country's flags"
	unknownMessage      = "Unknown command! Please write /start "
)

func (tg *telegramBot) printMessage(message *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	if text == commandStartMessage {
		tg.setHolidayRequest(message.Text)
		msg.ReplyMarkup = tg.countrySelection()
	}
	if err := tg.SendMessage(msg); err != nil {
		return errors.Wrap(err, "error with sending message")
	}
	return nil
}

func (tg *telegramBot) sendMessage(msg tgbotapi.MessageConfig) err {
	if _, err := tg.telegram.Send(msg); err != nil {
		return errors.Wrap(err, "send message failed")
	}
}
