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
	result := bs.db.Model(&model.Book{}).Count(&count)

	return count, result.Error
}

//TODO func (bs *BookStore) Books(offset, limit int) ([]model.Book, int64, error) {
  func (bs *BookStore) Books() ([]model.Book, error) {
	// var (
	// 	books []model.Book
	// 	count    int64
	// )

	// bs.db.Model(&books).Count(&count)
	// bs.db.Preload("Author")
	// bs.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&books)

	// return books, count, nil

  var books []model.Book

  result := bs.db.Preload("Author").Preload("Genre").Find(&books)

  return books, result.Error
}

func (bs *BookStore) NewBook(book *model.Book) error {

  result := bs.db.Create(book)

  return result.Error
}

func (bs *BookStore) GetBook(id uint64) (model.Book, error) {

  var book model.Book
  result := bs.db.First(&book, id)

  return book, result.Error
}
