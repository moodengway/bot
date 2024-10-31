package main

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/moodengway/bot/internal/bot"
	"github.com/moodengway/bot/internal/config"
	"github.com/moodengway/bot/internal/postgres"
	"github.com/moodengway/bot/internal/repository"
	"github.com/moodengway/bot/internal/service"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Panic("error loading config", zap.Error(err))
	}

	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		logger.Panic("error creating session", zap.Error(err))
	}

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		logger.Panic("error connecting postgres", zap.Error(err))
	}

	repo := repository.New(db)
	service := service.New(repo)

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
