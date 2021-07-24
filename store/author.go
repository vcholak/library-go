package store

import (
	"github.com.vcholak.library/model"
	"gorm.io/gorm"
)

type AuthorStore struct {
	db *gorm.DB
}

func NewAuthorStore(db *gorm.DB) *AuthorStore  {
	return &AuthorStore{
		db: db,
	}
}

func (bs *AuthorStore) AuthorCount() (int64, error) {
	var count int64
	bs.db.Model(&model.Author{}).Count(&count)

	return count, nil
}


