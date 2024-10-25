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
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{{
				Description: fmt.Sprintf("<@%s> requested to create a new match", i.Member.User.ID),
			}},
		},
	})
	if err != nil {
		b.logger.Error("error responding ack message", zap.Error(err))
		return
	}

	res, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		b.logger.Error("error getting ack message", zap.Error(err))
		return
	}

	match, err := b.service.CreateMatch(res.ID, i.Member.User.ID)
	if err != nil {
		b.logger.Error("error create a new match", zap.Error(err))
		return
	}

	matchEmbed := match.MessageEmbed()

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{&matchEmbed},
	})
	if err != nil {
		b.logger.Error("error editing interaction response", zap.Error(err))
	}
}
