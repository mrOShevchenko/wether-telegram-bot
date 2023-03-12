package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func (tg *TelegramBot) CallbackRequest(callback *tgbotapi.CallbackQuery) error {
	clbck := tgbotapi.NewCallback(callback.ID, callback.Data)

	if _, err := tg.bot.Request(clbck); err != nil {
		return errors.Wrap(err, "send request failed")
	}

	if clbck.Text != "" {
		if err := tg.PrintMessage(callback.Message, clbck.Text); err != nil {
			return errors.Wrap(err, "can't print message with actual holidays")
		}
	}
	return nil
}

func (tg *TelegramBot) CountrySelection(holidays *HolidayData) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.Australia.Location, tg.HolidayData.Australia.Flag),
				fmt.Sprintf("Holiday today: %s", holidays.Australia.Name)),
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.Ukraine.Location, tg.HolidayData.Ukraine.Flag),
				fmt.Sprintf("Holiday today: %s", holidays.Ukraine.Name)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.China.Location, tg.HolidayData.China.Flag),
				fmt.Sprintf("Holiday today: %s", holidays.China.Name)),
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.Canada.Location, tg.HolidayData.Canada.Flag),
				fmt.Sprintf("Holiday today: %s", holidays.Canada.Name)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.Georgia.Location, tg.HolidayData.Georgia.Flag),
				fmt.Sprintf("Holiday today: %s", holidays.Georgia.Name)),
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.France.Location, tg.HolidayData.France.Flag),
				fmt.Sprintf("Holiday today: %s", holidays.France.Name)),
		))
	return keyboard
}
