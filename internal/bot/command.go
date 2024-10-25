package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/util"
	"go.uber.org/zap"
)

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
