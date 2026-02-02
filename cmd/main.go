package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/freedom-sketch/sub2go-core/tg_bot/handlers"
	"github.com/go-telegram/bot"
)

func main() {
	err := config.Load("config.json")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
	}

	cfg := config.Get()

	b, err := bot.New(cfg.TelegramBot.Token, opts...)
	if err != nil {
		log.Panicf("Error creating bot instance: %v", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "key", bot.MatchTypeExact, handlers.Key)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back", bot.MatchTypeExact, handlers.Back)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "admin_panel", bot.MatchTypeExact, handlers.AdminPanel)

	b.Start(ctx)
}
