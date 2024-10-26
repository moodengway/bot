package bot

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

const (
	AcceptEmoji  string = "➕"
	Number1Emoji string = "1️⃣"
	Number2Emoji string = "2️⃣"
	Number3Emoji string = "3️⃣"
	Number4Emoji string = "4️⃣"
	Number5Emoji string = "5️⃣"
	Number6Emoji string = "6️⃣"
	Number7Emoji string = "7️⃣"
)

func (b *Bot) addReactionHandler() {
	reactionHandler := make(map[string]func(*discordgo.Session, *discordgo.MessageReactionAdd))
	reactionHandler[AcceptEmoji] = b.acceptReactionHandler()
	reactionHandler[Number1Emoji] = b.numberReactionHandler(1)
	reactionHandler[Number2Emoji] = b.numberReactionHandler(2)
	reactionHandler[Number3Emoji] = b.numberReactionHandler(3)
	reactionHandler[Number4Emoji] = b.numberReactionHandler(4)
	reactionHandler[Number5Emoji] = b.numberReactionHandler(5)
	reactionHandler[Number6Emoji] = b.numberReactionHandler(6)
	reactionHandler[Number7Emoji] = b.numberReactionHandler(7)

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

func (b *Bot) acceptReactionHandler() func(*discordgo.Session, *discordgo.MessageReactionAdd) {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
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

		clearEmoji(s, m.ChannelID, m.MessageID, b.logger)
		prepareNumberEmoji(s, m.ChannelID, m.MessageID, b.logger)
	}
}

func (b *Bot) numberReactionHandler(_ int) func(*discordgo.Session, *discordgo.MessageReactionAdd) {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		clearEmoji(s, m.ChannelID, m.MessageID, b.logger)
		prepareNumberEmoji(s, m.ChannelID, m.MessageID, b.logger)
	}
}

func clearEmoji(s *discordgo.Session, channelID string, messageID string, logger *zap.Logger) {
	err := s.MessageReactionsRemoveAll(channelID, messageID)
	if err != nil {
		logger.Warn("error clearing emoji", zap.Error(err))
	}
}

func prepareAcceptEmoji(s *discordgo.Session, channelID string, messageID string, logger *zap.Logger) {
	err := s.MessageReactionAdd(channelID, messageID, AcceptEmoji)
	if err != nil {
		logger.Warn("error adding accept reaction", zap.Error(err))
	}
}

func prepareNumberEmoji(s *discordgo.Session, channelID string, messageID string, logger *zap.Logger) {
	numbers := []string{Number1Emoji, Number2Emoji, Number3Emoji, Number4Emoji, Number5Emoji, Number6Emoji, Number7Emoji}
	for _, emoji := range numbers {
		err := s.MessageReactionAdd(channelID, messageID, emoji)
		if err != nil {
			logger.Warn("error adding number emoji", zap.Error(err), zap.String("emoji", emoji))
		}
	}
}
