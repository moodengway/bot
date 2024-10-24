package bot

import (
	"github.com/bwmarrin/discordgo"
)

type Session interface {
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

type Bot struct {
	session   Session
	channelID string
}

func New(channelID string, session Session) *Bot {
	return &Bot{
		session:   session,
		channelID: channelID,
	}
}

func (b *Bot) Send(message string) error {
	_, err := b.session.ChannelMessageSend(b.channelID, message)
	return err
}
