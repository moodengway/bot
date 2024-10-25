package service

import (
	"github.com/nonya123456/connect4/internal/model"
)

type Repository interface {
	CreateMatch(match model.Match) (model.Match, error)
	SaveMatch(match model.Match) (model.Match, error)
	FindMatchByMessageID(messageID string) (model.Match, bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateMatch(messageID string, host string) (model.Match, error) {
	match := model.Match{
		MessageID: messageID,
		Host:      host,
	}

	return s.repo.CreateMatch(match)
}

func (s *Service) AcceptMatch(messageID string, guest string) (model.Match, bool, error) {
	match, found, err := s.repo.FindMatchByMessageID(messageID)
	if err != nil {
		return model.Match{}, false, err
	}

	if !found || match.Guest != nil {
		return model.Match{}, false, nil
	}

	match.Guest = &guest

	match, err = s.repo.SaveMatch(match)
	if err != nil {
		return model.Match{}, false, err
	}

	return match, true, nil
}

func (s *Service) Place(i int) error {
	return nil
}
