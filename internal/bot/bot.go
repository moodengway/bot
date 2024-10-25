package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/service"
	"go.uber.org/zap"
)

type Bot struct {
	session *discordgo.Session
	service *service.Service
	logger  *zap.Logger
}

func New(session *discordgo.Session, service *service.Service, logger *zap.Logger) *Bot {
	return &Bot{
		session: session,
		service: service,
		logger:  logger,
	}
}

func (b *Bot) Start() error {
	err := b.session.Open()
	if err != nil {
		return err
	}

	_, err = b.session.ApplicationCommandCreate(b.session.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "create",
		Description: "Create new match",
	})
	if err != nil {
		return err
	}

	b.addCommandHandler()
	b.addReactionHandler()

	return nil
}

func (b *Bot) Stop() error {
	return b.session.Close()
}
