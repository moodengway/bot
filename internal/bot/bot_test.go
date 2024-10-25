package bot_test

import (
	"testing"

	"github.com/nonya123456/connect4/internal/bot"
	"github.com/nonya123456/connect4/internal/bot/mock"
	"github.com/stretchr/testify/suite"
)

type BotTestSuite struct {
	suite.Suite

	bot       *bot.Bot
	channelID string

	mockSession *mock.MockSession
}

func TestBotTestSuite(t *testing.T) {
	suite.Run(t, new(BotTestSuite))
}

func (t *BotTestSuite) SetupTest() {
	channelID := "1234"
	mockSession := new(mock.MockSession)

	t.bot = bot.New(channelID, mockSession)
	t.mockSession = mockSession
	t.channelID = channelID
}

func (t *BotTestSuite) TestSend() {
	message := "Test message"
	t.mockSession.On("ChannelMessageSend", t.channelID, message).Return(nil, nil).Once()

	err := t.bot.Send(message)
	t.NoError(err)
}
