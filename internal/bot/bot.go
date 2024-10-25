package bot

import (
	"fmt"

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

	commandHandlers := make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate))
	commandHandlers["create"] = b.createCommandHandler

	b.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	return nil
}

func (b *Bot) Stop() error {
	return b.session.Close()
}

func (b *Bot) createCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Member.User.ID

	match, err := b.service.CreateMatch(userID)
	if err != nil {
		b.logger.Error("error create a new match", zap.Error(err))
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Created match#%d", match.ID),
		},
	})

	if err != nil {
		b.logger.Error("error responding created", zap.Error(err))
	}
}
