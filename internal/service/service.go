package service

import (
	"github.com/nonya123456/connect4/internal/model"
)

type Repository interface {
	CreateMatch(messageID string, host string) (model.Match, error)
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
	return s.repo.CreateMatch(messageID, host)
}

func (s *Service) Place(i int) error {
	return nil
}
