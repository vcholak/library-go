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

func (as *AuthorStore) AuthorCount() (int64, error) {
	var count int64
	result := as.db.Model(&model.Author{}).Count(&count)

	return count, result.Error
}

func (as *AuthorStore) NewAuthor(author *model.Author) error {

  result := as.db.Create(author)

  return result.Error
}

func (as *AuthorStore) Authors() ([]model.Author, error) {

  var authors []model.Author

  result := as.db.Find(&authors)

  return authors, result.Error
}

func (as *AuthorStore) GetAuthor(id uint64) (model.Author, error) {

  var author model.Author
  result := as.db.Preload("Books").First(&author, id)

  return author, result.Error
}
