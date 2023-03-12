package main

import (
	"context"
	"log"
	"task2.3.3/internal"
	"task2.3.3/internal/telegram"
)

//TODO: дописать телеграм клиент интерфейсами

func main() {
	ctx := context.Background()
	container, err := internal.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, "log", container.NewLogger())
	tg, err := telegram.New(container)
	if err != nil {
		log.Fatalf("error in creating new tgBot: %v", err)
	}
	tg.Running()

}
