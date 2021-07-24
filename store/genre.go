package store

import (
	"github.com.vcholak.library/model"
	"gorm.io/gorm"
)

type GenreStore struct {
	db *gorm.DB
}

func NewGerneStore(db *gorm.DB) *GenreStore  {
	return &GenreStore{
		db: db,
	}
}

func (bs *GenreStore) GenreCount() (int64, error) {
	var count int64
	bs.db.Model(&model.Genre{}).Count(&count)

	return count, nil
}
