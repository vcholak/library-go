package store

import (
	"github.com.vcholak.library/copy"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookInstanceStore struct {
	db *gorm.DB
}

func NewBookInstanceStore(db *gorm.DB) *BookInstanceStore {
	return &BookInstanceStore{
		db: db,
	}
}

func (bs *BookInstanceStore) BookInstanceCount() (int64, error) {
	var count int64
	bs.db.Model(&copy.BookInstance{}).Count(&count)

	return count, nil
}

func (bs *BookInstanceStore) AvailableBookInstanceCount() (int64, error) {
	var count int64
	bs.db.Model(&copy.BookInstance{}).Where("Status=?", 5).Count(&count)

	return count, nil
}

func (bs *BookInstanceStore) BookInstances() ([]copy.BookInstance, error) {

	var copies []copy.BookInstance

	result := bs.db.Preload("Book").Find(&copies)

	return copies, result.Error
}

func (bs *BookInstanceStore) NewBookInstance(copy *copy.BookInstance) error {

	result := bs.db.Preload(clause.Associations).Create(copy)

	return result.Error
}

func (bs *BookInstanceStore) GetBookInstance(id uint64) (copy.BookInstance, error) {

	var bookCopy copy.BookInstance
	result := bs.db.Preload("Book").Preload("Book.Author").First(&bookCopy, id)

	return bookCopy, result.Error
}

func (bs *BookInstanceStore) UpdateBookInstance(bookCopy *copy.BookInstance) error {

	result := bs.db.Save(bookCopy)

	return result.Error
}
