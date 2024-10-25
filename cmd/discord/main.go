package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/nonya123456/connect4/internal/bot"
	"github.com/nonya123456/connect4/internal/config"
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

	bot := bot.New(cfg.ChannelID, session)

	if err = bot.Start(); err != nil {
		logger.Panic("error opening connection", zap.Error(err))
	}
	defer func() {
		_ = bot.Stop()
	}()

	if err = bot.Send("Hello, World"); err != nil {
		logger.Warn("error sending message", zap.Error(err))
	}
}
