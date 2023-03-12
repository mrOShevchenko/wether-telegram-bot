package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func (tg *TelegramBot) OnCommandCreate(message *tgbotapi.Message) error {
	switch message.Text {
	case "/start":
		if err := tg.PrintMessage(message, commandStartMessage); err != nil {
			return errors.Wrap(err, "error with case /start ")
		}

	default:
		if err := tg.PrintMessage(message, greetingMessage); err != nil {
			return errors.Wrap(err, "can't print greetingMessage")
		}
	}
	return nil
}

func (tg *TelegramBot) OnCallbackQuery(callback *tgbotapi.CallbackQuery) error {
	if err := tg.CallbackRequest(callback); err != nil {
		return errors.Wrap(err, "error in onCallbackQuery: ")
	}
	return nil
}
