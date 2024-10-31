package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/moodengway/bot/internal/util"
	"go.uber.org/zap"
)

func (b *Bot) addCommandHandler() {
	commandHandlers := make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate))
	commandHandlers["create"] = b.createCommandHandler()

	b.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func (b *Bot) createCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

		prepareEmoji(s, i.ChannelID, res.ID, b.logger)
	}
}

func prepareEmoji(s *discordgo.Session, channelID string, messageID string, logger *zap.Logger) {
	err := s.MessageReactionAdd(channelID, messageID, AcceptEmoji)
	if err != nil {
		logger.Warn("error adding accept reaction", zap.Error(err))
	}

	numbers := []string{Number1Emoji, Number2Emoji, Number3Emoji, Number4Emoji, Number5Emoji, Number6Emoji, Number7Emoji}
	for _, emoji := range numbers {
		err := s.MessageReactionAdd(channelID, messageID, emoji)
		if err != nil {
			logger.Warn("error adding number emoji", zap.Error(err), zap.String("emoji", emoji))
		}
	}
}
