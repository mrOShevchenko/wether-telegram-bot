package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Telegram
	AbstractAPI
}

type Telegram struct {
	TokenENV    string `split_words:"true"`
	APIBaseURL  string `split_words:"true"`
	SendMessage string `split_words:"true"`
}

type AbstractAPI struct {
	AbstractApiBaseUrl  string `envconfig:"ABSTRACT_API_BASE_URL"`
	AbstractApiTokenEnv string `envconfig:"ABSTRACT_API_TOKEN_ENV"`
}

func Get() (*Config, error) {
	CfgTG, err := ReadConfigTG("env/telegramBot.env")
	if err != nil {
		log.Fatalf("can't read config from file, error: %w", err)
		return nil, err
	}
	CfgAbstract, err := ReadConfigAbstractAPI("/home/shevchenko/GolandProjects/education/task2.3.3/env/abstractAPI.env")
	if err != nil {
		log.Fatalf("can't read config from file, error: %v", err)
		return nil, err
	}
	cfg := &Config{
		Telegram:    *CfgTG,
		AbstractAPI: *CfgAbstract,
	}
	return cfg, nil
}

func ReadConfigTG(pathToEnv string) (*Telegram, error) {
	if err := godotenv.Load(pathToEnv); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}
	c := &Telegram{}

	if err := envconfig.Process("TELEGRAM_API", c); err != nil {
		return nil, fmt.Errorf("error process config: %w", err)
	}

	return c, nil
}

func ReadConfigAbstractAPI(pathToEnv string) (*AbstractAPI, error) {
	if err := godotenv.Load(pathToEnv); err != nil {
		return nil, fmt.Errorf("error parsing config: %v", err)
	}
	c := &AbstractAPI{}

	if err := envconfig.Process("", c); err != nil {
		return nil, fmt.Errorf("error process config: %w", err)
	}
	return c, nil
}
