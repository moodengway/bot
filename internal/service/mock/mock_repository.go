package mock

import (
	"github.com/moodengway/bot/internal/model"
	"github.com/moodengway/bot/internal/service"
	"github.com/stretchr/testify/mock"
)

var _ service.Repository = (*MockRepository)(nil)

type MockRepository struct {
	mock.Mock
}

// FindMatchByMessageID implements service.Repository.
func (m *MockRepository) FindMatchByMessageID(messageID string) (model.Match, bool, error) {
	args := m.Called(messageID)

	match, ok := args.Get(0).(model.Match)
	if !ok {
		panic("cannot convert to match")
	}

	return match, args.Bool(1), args.Error(2)
}

// SaveMatch implements service.Repository.
func (m *MockRepository) SaveMatch(match model.Match) (model.Match, error) {
	args := m.Called(match)

	match, ok := args.Get(0).(model.Match)
	if !ok {
		panic("cannot convert to match")
	}

	return match, args.Error(1)
}

// CreateMatch implements service.Repository.
func (m *MockRepository) CreateMatch(match model.Match) (model.Match, error) {
	args := m.Called(match)

	match, ok := args.Get(0).(model.Match)
	if !ok {
		panic("cannot convert to match")
	}

	return match, args.Error(1)
}
