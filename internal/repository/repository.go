package repository

import (
	"errors"

	"github.com/nonya123456/connect4/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateMatch(match model.Match) (model.Match, error) {
	if err := r.db.Create(&match).Error; err != nil {
		return model.Match{}, err
	}

	return match, nil
}

func (r *Repository) SaveMatch(match model.Match) (model.Match, error) {
	if err := r.db.Save(&match).Error; err != nil {
		return model.Match{}, err
	}

	return match, nil
}

func (r *Repository) FindMatchByMessageID(messageID string) (model.Match, bool, error) {
	var match model.Match
	if err := r.db.Where("message_id = ?", messageID).First(&match).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Match{}, false, nil
		}

		return model.Match{}, false, err
	}

	return match, true, nil
}
