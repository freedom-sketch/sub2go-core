package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/freedom-sketch/sub2go-core/infra/logger"
	hdls "github.com/freedom-sketch/sub2go-core/tg_bot/handlers"
	"github.com/go-telegram/bot"
)

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}

	logger, err := logger.NewLogger("logs/tg_bot.log", "info")
	if err != nil {
		log.Panicf("Failed to create logger: %v:", err)
	}
	defer logger.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(hdls.DefaultHandler),
	}

	b, err := bot.New(cfg.TelegramBot.Token, opts...)
	if err != nil {
		logger.Errorf("Error creating bot instance: %v", err)
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypePrefix, hdls.StartHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "key", bot.MatchTypeExact, hdls.Key)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back", bot.MatchTypeExact, hdls.Back)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "admin_panel", bot.MatchTypeExact, hdls.AdminPanel)

	b.Start(ctx)
}
