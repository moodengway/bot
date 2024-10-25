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
		logger.Panic("load env failed", zap.Error(err))
	}

	cfg := config.New()

	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		logger.Panic("create session failed", zap.Error(err))
	}
	defer func() {
		_ = session.Close()
	}()

	bot := bot.New(cfg.ChannelID, session)
	if err = bot.Send("Hello, World"); err != nil {
		logger.Warn("send message failed", zap.Error(err))
	}
}
