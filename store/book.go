package store

import (
	"github.com.vcholak.library/model"

	"gorm.io/gorm"
)

type BookStore struct {
	db *gorm.DB
}

func NewBookStore(db *gorm.DB) *BookStore  {
	return &BookStore{
		db: db,
	}	
}

func (bs *BookStore) BookCount() (int64, error) {
	var count int64
	bs.db.Model(&model.Book{}).Count(&count)

	return count, nil
}

func (bs *BookStore) Books(offset, limit int) ([]model.Book, int64, error) {
	var (
		books []model.Book
		count    int64
	)

	bs.db.Model(&books).Count(&count)
	//bs.db.Preload("Favorites"). - for associations
	//	Preload("Tags").
	//	Preload("Author").
	bs.db.Offset(offset).
		Limit(limit).
		Order("created_at desc").Find(&books)

	return books, count, nil
}

