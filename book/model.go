package book

import (
	"github.com.vcholak.library/genre"

	"gorm.io/datatypes"
)

// `Book` belongs to `Author`, `AuthorID` is the foreign key
// `Book` belongs to `Genre`, `GenreID` is the foreign key
type Book struct {
	ID       uint        `json:"id"`
	Title    string      `gorm:"unique_index;not null" json:"title"`
	AuthorId uint        `json:"author_id"`
	Author   Author      `json:"author"` // `gorm:"foreignkey:AuthorID"`
	Summary  string      `gorm:"not null" json:"summary"`
	ISBN     string      `gorm:"not null" json:"isbn"`
	GenreId  uint        `json:"genre_id"`
	Genre    genre.Genre `json:"genre"` // `gorm:"foreignkey:ID"`
}

type Author struct {
	ID         uint           `json:"id"`
	FirstName  string         `gorm:"not null" json:"first_name"`
	FamilyName string         `gorm:"not null" json:"family_name"`
	BirthDate  datatypes.Date `gorm:"not null" json:"birth_date"`
	DeathDate  datatypes.Date `json:"death_date"`
	LifeSpan   string         `gorm:"not null" json:"life_span"`
	Books      []Book         `json:"books"`
}
