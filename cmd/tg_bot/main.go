package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	hdls "github.com/freedom-sketch/sub2go-core/tg_bot/handlers"
	"github.com/go-telegram/bot"
)

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(hdls.DefaultHandler),
	}

	b, err := bot.New(cfg.TelegramBot.Token, opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, hdls.StartHandler)

	b.Start(ctx)
}
