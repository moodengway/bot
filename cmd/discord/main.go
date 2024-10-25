package main

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/nonya123456/connect4/internal/bot"
	"github.com/nonya123456/connect4/internal/config"
	"github.com/nonya123456/connect4/internal/service"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	err := godotenv.Load()
	if err != nil {
		logger.Panic("error loading env file", zap.Error(err))
	}

	cfg := config.New()

	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		logger.Panic("error creating session", zap.Error(err))
	}

	service := service.New()

	bot := bot.New(session, service, logger)

	if err = bot.Start(); err != nil {
		logger.Panic("error opening connection", zap.Error(err))
	}
	defer func() {
		_ = bot.Stop()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	logger.Info("bot is now running")
	<-stop
}
