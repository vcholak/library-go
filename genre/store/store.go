package store

import (
	"github.com.vcholak.library/genre"

	"gorm.io/gorm"
)

type GenreStore struct {
	db *gorm.DB
}

func NewGenreStore(db *gorm.DB) *GenreStore {
	return &GenreStore{
		db: db,
	}
}

func (gs *GenreStore) GenreCount() (int64, error) {
	var count int64
	gs.db.Model(&genre.Genre{}).Count(&count)

	return count, nil
}

func (gs *GenreStore) NewGenre(genre *genre.Genre) error {

	result := gs.db.Create(genre)

	return result.Error
}

func (gs *GenreStore) Genres() ([]genre.Genre, error) {

	var genres []genre.Genre

	result := gs.db.Find(&genres)

	return genres, result.Error
}

func (gs *GenreStore) GetGerne(id uint64) (genre.Genre, error) {

	var gerne genre.Genre
	result := gs.db.First(&gerne, id)

	return gerne, result.Error
}

func (gs *GenreStore) UpdateGenre(genre *genre.Genre) error {

	result := gs.db.Save(genre)

	return result.Error
}
