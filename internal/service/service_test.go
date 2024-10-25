package service_test

import (
	"testing"

	"github.com/nonya123456/connect4/internal/model"
	"github.com/nonya123456/connect4/internal/service"
	"github.com/nonya123456/connect4/internal/service/mock"
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

	t.mockRepo.On("CreateMatch", messageID, host).
		Return(model.Match{ID: matchID, MessageID: messageID, Host: host}, nil).
		Once()

	match, err := t.underTest.CreateMatch(messageID, host)
	t.NoError(err)

	t.Equal(model.Match{ID: matchID, MessageID: messageID, Host: host}, match)
}
