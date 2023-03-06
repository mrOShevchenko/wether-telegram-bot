package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func (tg *telegramBot) callbackRequest(callback *tgbotapi.CallbackQuery) error {
	clbck := tgbotapi.NewCallback(callback.ID, callback.Data)
	if _, err := tg.bot.Request(clbck); err != nil {
		return errors.Wrap(err, "send request failed")
	}
	if clbck.Text != "" {
		holidays, err := tg.holidayTodayRequest{clbck.Text}
		if err != nil {
			return errors.Wrap(err, "can't get holidayTodayRequest")
		}
		if err := tg.printMessage(callback.Message, holidays); err != nil {
			return errors.Wrap(err, "can't print message with actual holidays"+
				"")
		}
	}
	return nil
}

func (tg *telegramBot) countrySelection() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %v", tg.holidayData.Australia.Location, tg.holidayData.Australia.Flag),
				fmt.Sprintf("Holiday today: %d - %s", tg.holidayData.Today, tg.holidayData.Australia.Holiday)),
			tgbotapi.NewInlineKeyboardButtonData("Ukraine-UA🇺🇦", countryData.Ukraine.Name),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("China-CN🇨🇳", countryData.China.Name),
			tgbotapi.NewInlineKeyboardButtonData("Canada-CA🇨🇦", countryData.Canada.Name),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Georgia-GE🇬🇪", countryData.Georgia.Name),
			tgbotapi.NewInlineKeyboardButtonData("France-FR🇫🇷", countryData.France.Name),
		))
	return keyboard
}
