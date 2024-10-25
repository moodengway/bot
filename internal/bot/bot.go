package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/service"
	"github.com/nonya123456/connect4/internal/util"
	"go.uber.org/zap"
)

const (
	AcceptEmoji string = "➕"
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

func (b *Bot) addCommandHandler() {
	commandHandlers := make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate))
	commandHandlers["create"] = b.createCommandHandler

	b.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func (b *Bot) createCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{{
				Description: fmt.Sprintf("Creating a new match from %s", util.Mention(i.Member.User.ID)),
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
		return
	}

	err = s.MessageReactionAdd(i.ChannelID, res.ID, AcceptEmoji)
	if err != nil {
		b.logger.Warn("error adding accept reaction", zap.Error(err))
	}
}

func (b *Bot) addReactionHandler() {
	reactionHandler := make(map[string]func(*discordgo.Session, *discordgo.MessageReactionAdd))
	reactionHandler[AcceptEmoji] = b.acceptReactionHandler

	b.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		if m.UserID == s.State.User.ID {
			return
		}

		h, ok := reactionHandler[m.Emoji.Name]
		if !ok {
			return
		}

		message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
		if err != nil {
			b.logger.Error("error getting reacted message", zap.Error(err))
			return
		}

		if message.Author.ID != s.State.User.ID {
			return
		}

		h(s, m)
	})
}

func (b *Bot) acceptReactionHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	match, ok, err := b.service.AcceptMatch(m.MessageID, m.UserID)
	if err != nil {
		b.logger.Error("error accepting match", zap.Error(err))
		return
	}

	if !ok {
		b.logger.Debug("match is not found or is already accepted")
		return
	}

	matchEmbed := match.MessageEmbed()
	_, err = s.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, &matchEmbed)
	if err != nil {
		b.logger.Error("error editing embed", zap.Error(err))
		return
	}

	err = s.MessageReactionsRemoveAll(m.ChannelID, m.MessageID)
	if err != nil {
		b.logger.Warn("error removing reactions", zap.Error(err))
	}
}
