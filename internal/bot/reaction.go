package bot

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

const (
	AcceptEmoji string = "➕"
)

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
