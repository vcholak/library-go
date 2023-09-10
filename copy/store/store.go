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
	bs.db.Model(&copy.BookInstance{}).Count(&count) //TODO fix

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
