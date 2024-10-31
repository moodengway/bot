package service_test

import (
	"testing"

	"github.com/moodengway/bot/internal/model"
	"github.com/moodengway/bot/internal/service"
	"github.com/moodengway/bot/internal/service/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite

	underTest *service.Service

	mockRepo *mock.MockRepository
}

func TestBotTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (t *ServiceTestSuite) SetupTest() {
	mockRepo := new(mock.MockRepository)

	t.underTest = service.New(mockRepo)
	t.mockRepo = mockRepo
}

func (t *ServiceTestSuite) TestCreateMatch() {
	var matchID uint = 1
	var messageID string = "testmessageid"
	var host string = "testhost"
	var boardString string = "000000000000000000000000000000000000000000"
	var roundNumber int = 1

	expectedInputMatch := model.Match{
		MessageID:   messageID,
		Host:        host,
		BoardString: boardString,
		RoundNumber: roundNumber,
	}

	mockOutputMatch := model.Match{
		ID:          matchID,
		MessageID:   messageID,
		Host:        host,
		BoardString: boardString,
		RoundNumber: roundNumber,
	}

	t.mockRepo.On("CreateMatch", expectedInputMatch).Return(mockOutputMatch, nil).Once()

	match, err := t.underTest.CreateMatch(messageID, host)
	t.NoError(err)
	t.Equal(mockOutputMatch, match)
}

func (t *ServiceTestSuite) TestAcceptMatchNotFound() {
	var messageID string = "testmessageid"
	var guest string = "testguest"

	t.mockRepo.On("FindMatchByMessageID", messageID).Return(model.Match{}, false, nil).Once()

	_, ok, err := t.underTest.AcceptMatch(messageID, guest)
	t.NoError(err)
	t.Equal(false, ok)
}

func (t *ServiceTestSuite) TestAcceptMatchAlreadyAccepted() {
	var messageID string = "testmessageid"
	var guest string = "testguest"

	t.mockRepo.On("FindMatchByMessageID", messageID).Return(model.Match{Guest: &guest}, true, nil).Once()

	_, ok, err := t.underTest.AcceptMatch(messageID, guest)
	t.NoError(err)
	t.Equal(false, ok)
}

func (t *ServiceTestSuite) TestAcceptMatch() {
	var id uint = 1
	var messageID string = "testmessageid"
	var host string = "testhost"
	var guest string = "testguest"

	match := model.Match{
		ID:        id,
		MessageID: messageID,
		Host:      host,
	}

	t.mockRepo.On("FindMatchByMessageID", messageID).Return(match, true, nil).Once()

	match.Guest = &guest

	t.mockRepo.On("SaveMatch", match).Return(match, nil).Once()

	updatedMatch, ok, err := t.underTest.AcceptMatch(messageID, guest)
	t.NoError(err)
	t.Equal(true, ok)
	t.Equal(model.Match{ID: id, MessageID: messageID, Host: host, Guest: &guest}, updatedMatch)
}
