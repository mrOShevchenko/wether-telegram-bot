package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func (tg *telegramBot) callbackRequest(callback *tgbotapi.CallbackQuery) error {
	clbck := tgbotapi.NewCallback(callback.ID, callback.Data)
	fmt.Printf("\n\n clbck %v\n\n", clbck)

	if _, err := tg.bot.Request(clbck); err != nil {
		return errors.Wrap(err, "send request failed")
	}
	fmt.Printf("\n\n clbck 22 %v\n\n", clbck)

	if clbck.Text != "" {
		//actualHoliday, err := tg.holidayRequestforCountry(clbck.Text)
		//if err != nil {
		//	return errors.Wrapf(err, "error in holidayRequestforCountry %+w")
		//}

		if err := tg.printMessage(callback.Message, actualHoliday); err != nil {
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
				fmt.Sprintf("%s %s", tg.HolidayData.Australia.Location, tg.HolidayData.Australia.Flag),
				fmt.Sprintf("Holiday today: %s", tg.HolidayData.Australia.Name)),
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.Ukraine.Location, tg.HolidayData.Ukraine.Flag),
				fmt.Sprintf("Holiday today: %s", tg.HolidayData.Ukraine.Name)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.China.Location, tg.HolidayData.China.Flag),
				fmt.Sprintf("Holiday today: %s", tg.HolidayData.China.Name)),
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.Canada.Location, tg.HolidayData.Canada.Flag),
				fmt.Sprintf("Holiday today: %s", tg.HolidayData.Canada.Name)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.Georgia.Location, tg.HolidayData.Georgia.Flag),
				fmt.Sprintf("Holiday today: %s", tg.HolidayData.Georgia.Name)),
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s %s", tg.HolidayData.France.Location, tg.HolidayData.France.Flag),
				fmt.Sprintf("Holiday today: %s", tg.HolidayData.France.Name)),
		))
	return keyboard
}
