package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

const (
	commandStartMessage = "OK, let's see what holiday today at this countries:"
	greetingMessage     = " Hello I am Name Bot. Press /start to see country's flags"
	unknownMessage      = "Unknown command! Please write /start "
)

func (tg *telegramBot) printMessage(message *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	if text == commandStartMessage {
		holidays, err := tg.holidayRequest()
		if err != nil {
			return errors.Wrapf(err, "error in holidayRequest %+v")
		}
		fmt.Printf("\n\n holidays in HolidayRequest when was commandStartMessage was selected %v", holidays)

		msg.ReplyMarkup = tg.countrySelection()
	}
	if err := tg.sendMessage(msg); err != nil {
		return errors.Wrap(err, "error with sending message")
	}
	return nil
}

func (tg *telegramBot) sendMessage(msg tgbotapi.MessageConfig) error {
	if _, err := tg.bot.Send(msg); err != nil {
		return errors.Wrap(err, "send message failed")
	}
	return nil
}
