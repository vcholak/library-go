package store

import (
	"github.com.vcholak.library/model"
	"gorm.io/gorm"
)

type BookInstanceStore struct {
	db *gorm.DB
}

func NewBookInstanceStore(db *gorm.DB) *BookInstanceStore  {
	return &BookInstanceStore{
		db: db,
	}
}

func (bs *BookInstanceStore) BookInstanceCount() (int64, error) {
	var count int64
	bs.db.Model(&model.BookInstance{}).Count(&count)

	return count, nil
}

func (bs *BookInstanceStore) AvailableBookInstanceCount() (int64, error) {
	var count int64
	bs.db.Model(&model.BookInstance{}).Count(&count) //TODO fix

	return count, nil
}
