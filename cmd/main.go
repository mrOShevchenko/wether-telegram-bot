package main

import (
	"context"
	"log"
	"task2.3.3/internal"
	"task2.3.3/internal/telegram"
)

//TODO: url - разобраться почему не обновляется страна (дата и все остальное - отлично)
//TODO: дописать телеграм клиент интерфейсами
//
//TODO: поставить отарпвку сообщений в чат, при ответе, а не только принт

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
