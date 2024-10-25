package mock

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/connect4/internal/bot"
	"github.com/stretchr/testify/mock"
)

var _ bot.Session = (*MockSession)(nil)

type MockSession struct {
	mock.Mock
}

func (m *MockSession) Open() error {
	panic("unimplemented")
}

func (m *MockSession) Close() error {
	panic("unimplemented")
}

func (m *MockSession) ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error) {
	var args mock.Arguments

	if len(options) > 0 {
		args = m.Called(channelID, content, options)
	} else {
		args = m.Called(channelID, content)
	}

	message, ok := args.Get(0).(*discordgo.Message)
	if !ok {
		return nil, args.Error(1)
	}

	return message, args.Error(1)
}
