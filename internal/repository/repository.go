package repository

import (
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

func (r *Repository) CreateMatch(host string) (model.Match, error) {
	match := model.Match{
		Host: host,
	}

	if err := r.db.Create(&match).Error; err != nil {
		return model.Match{}, err
	}

	return match, nil
}
