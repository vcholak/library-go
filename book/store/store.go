package store

import (
	"github.com.vcholak.library/book"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	result := bs.db.Model(&book.Book{}).Count(&count)

	return count, result.Error
}

//TODO func (bs *BookStore) Books(offset, limit int) ([]model.Book, int64, error) {
  func (bs *BookStore) Books() ([]book.Book, error) {
	// var (
	// 	books []model.Book
	// 	count    int64
	// )

	// bs.db.Model(&books).Count(&count)
	// bs.db.Preload("Author")
	// bs.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&books)

	// return books, count, nil

  var books []book.Book

  result := bs.db.Preload("Author").Preload("Genre").Find(&books)

  return books, result.Error
}

func (bs *BookStore) NewBook(book *book.Book) error {

  result := bs.db.Preload(clause.Associations).Create(book)

  return result.Error
}

func (bs *BookStore) GetBook(id uint64) (book.Book, error) {

  var book book.Book
  result := bs.db.Preload("Author").Preload("Genre").First(&book, id)

  return book, result.Error
}

//------------------------------
type AuthorStore struct {
	db *gorm.DB
}

func NewAuthorStore(db *gorm.DB) *AuthorStore  {
	return &AuthorStore{
		db: db,
	}
}

func (as *AuthorStore) AuthorCount() (int64, error) {
	var count int64
	result := as.db.Model(&book.Author{}).Count(&count)

	return count, result.Error
}

func (as *AuthorStore) NewAuthor(author *book.Author) error {

  result := as.db.Create(author)

  return result.Error
}

func (as *AuthorStore) Authors() ([]book.Author, error) {

  var authors []book.Author

  result := as.db.Find(&authors)

  return authors, result.Error
}

func (as *AuthorStore) GetAuthor(id uint64) (book.Author, error) {

  var author book.Author
  result := as.db.Preload("Books").First(&author, id)

  return author, result.Error
}
