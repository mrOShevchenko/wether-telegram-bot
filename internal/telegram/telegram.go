package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"net/http"
	"task2.3.3/internal"
	"task2.3.3/internal/config"
	"time"
)

type TgClient interface {
	Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
	Send(chattable tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}

type TelegramBot struct {
	c           internal.Container
	bot         TgClient
	updatesCh   tgbotapi.UpdatesChannel
	ctx         context.Context
	HolidayData *HolidayData
}

func (tg *TelegramBot) Running() {
	for update := range tg.updatesCh {
		tg.EventUpdates(update)
	}
}

func (tg *TelegramBot) EventUpdates(update tgbotapi.Update) {
	log := tg.c.NewLogger()
	switch {
	case update.CallbackQuery != nil:
		if err := tg.OnCallbackQuery(update.CallbackQuery); err != nil {
			log.Errorf("error in callback: %v", err)
		}
	case update.Message != nil:
		if err := tg.OnCommandCreate(update.Message); err != nil {
			log.Errorf("error woth command : %v", err)
		}
	default:
		log.Infof("unkwown event: %+v\n", update)
	}
}

func New(c internal.Container) (*TelegramBot, error) {
	log := c.NewLogger()
	cfg := c.NewConfig()

	bot, err := newBot(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can't create new Bot")
	}

	actualHolidays := newHoliday()
	t := &TelegramBot{
		c:           c,
		bot:         bot,
		HolidayData: actualHolidays,
	}

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
