package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"net/http"
	"task2.3.3/internal"
	"task2.3.3/internal/config"
	"time"
)

type telegramBot struct {
	c           internal.Container
	bot         *tgbotapi.BotAPI
	updatesCh   tgbotapi.UpdatesChannel
	ctx         context.Context
	HolidayData *HolidayData
}

func (tg *telegramBot) Running() {
	for update := range tg.updatesCh {
		tg.eventUpdates(update)
	}
}

func (tg *telegramBot) eventUpdates(update tgbotapi.Update) {
	log := tg.c.NewLogger()
	switch {
	case update.CallbackQuery != nil:
		fmt.Printf("\n\n\ncase update.CallbackQuery DATA: %v\n\n\n", update.CallbackQuery.Data)
		if err := tg.onCallbackQuery(update.CallbackQuery); err != nil {
			log.Errorf("error in callback: %v", err)
		}
	case update.Message != nil:
		fmt.Printf("\n\n\ncase update.Message: %v\n\n\n", update.Message)
		if err := tg.onCommandCreate(update.Message); err != nil {
			log.Errorf("error woth command : %v", err)
		}
	default:
		log.Infof("unkwown event: %+v\n", update)
	}
}

func New(c internal.Container) (*telegramBot, error) {
	log := c.NewLogger()
	cfg := c.NewConfig()

	bot, err := newBot(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can't create new Bot")
	}

	actualHolidays := newHoliday()
	t := &telegramBot{
		c:           c,
		bot:         bot,
		HolidayData: actualHolidays,
	}
	actualHolidays, err = t.holidayRequest()
	if err != nil {
		return nil, errors.Wrapf(err, "error in actualHoliday - holidayRequest %w")
	}
	fmt.Printf("\n/\n/\nactualHolidays in new Bot : %s\n/\n/\n", actualHolidays)
	botUpdate := tgbotapi.NewUpdate(0)
	botUpdate.Timeout = 60
	t.updatesCh = t.bot.GetUpdatesChan(botUpdate)

	log.Infof("Authorized on account %s", bot.Self.UserName)
	return t, nil
}

func newBot(cfg *config.Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(cfg.TokenENV, tgbotapi.APIEndpoint, &http.Client{Timeout: 60 * time.Second})
	if err != nil {
		return nil, errors.Wrap(err, "tgbotapi.NewBotAPIWithClient() failed with error:")
	}
	bot.Debug = true
	return bot, nil
}
