package store

import (
	"fmt"

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

func (gs *GenreStore) GenreCount() (int64, error) {
	var count int64
	gs.db.Model(&model.Genre{}).Count(&count)

	return count, nil
}

func (gs *GenreStore) NewGenre(genre *model.Genre) {

  result := gs.db.Create(genre)

  if result.Error != nil {
    fmt.Println(result.Error)
    panic("Failed to create a new genre")
  }
}

func (gs *GenreStore) Genres() ([]model.Genre, error) {

  var genres []model.Genre

  result := gs.db.Find(&genres)

  return genres, result.Error
}

func (gs *GenreStore) GetGerne(id uint64) (model.Genre, error) {

  var gerne model.Genre
  result := gs.db.First(&gerne, id)

  return gerne, result.Error
}
