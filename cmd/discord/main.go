package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/config"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	err := config.LoadENV()
	if err != nil {
		logger.Panic("load env failed", zap.Error(err))
	}

	cfg := config.New()

	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		logger.Panic("create session failed", zap.Error(err))
	}

	_, err = discord.ChannelMessageSend(cfg.ChannelID, "Hello, World")
	if err != nil {
		logger.Warn("send message failed", zap.Error(err))
	}
}
