package service

import (
	"github.com/nonya123456/connect4/internal/model"
	"github.com/nonya123456/connect4/internal/util"
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
	var board [6][7]int

	match := model.Match{
		MessageID:   messageID,
		Host:        host,
		BoardString: util.BoardString(board),
		RoundNumber: 1,
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

// TODO: add test
func (s *Service) Place(messageID string, userID string, i int) (model.Match, bool, error) {
	match, found, err := s.repo.FindMatchByMessageID(messageID)
	if err != nil {
		return model.Match{}, false, err
	}

	if !found || match.Guest == nil {
		return model.Match{}, false, nil
	}

	if match.RoundNumber%2 == 1 && match.Host != userID {
		return model.Match{}, false, nil
	}

	if match.RoundNumber%2 == 0 && *match.Guest != userID {
		return model.Match{}, false, nil
	}

	board, err := match.Board()
	if err != nil {
		return model.Match{}, false, err
	}

	placed := false
	for j := 0; j < 6; j++ {
		if board[j][i-1] != 0 {
			continue
		}

		board[j][i-1] = 2 - (match.RoundNumber % 2)
		placed = true
		break
	}

	if !placed {
		return model.Match{}, false, nil
	}

	match.BoardString = util.BoardString(board)
	match.RoundNumber += 1

	match, err = s.repo.SaveMatch(match)
	if err != nil {
		return model.Match{}, false, err
	}

	return match, true, nil
}
