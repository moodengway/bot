package mock

import (
	"github.com/nonya123456/connect4/internal/model"
	"github.com/nonya123456/connect4/internal/service"
	"github.com/stretchr/testify/mock"
)

var _ service.Repository = (*MockRepository)(nil)

type MockRepository struct {
	mock.Mock
}

// CreateMatch implements service.Repository.
func (m *MockRepository) CreateMatch(messageID string, host string) (model.Match, error) {
	args := m.Called(messageID, host)

	match, ok := args.Get(0).(model.Match)
	if !ok {
		panic("cannot convert to match")
	}

	return match, args.Error(1)
}
