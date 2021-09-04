package model

import (
	"gorm.io/datatypes"
)

type Author struct {
  ID         uint `json:"id"`
	FirstName  string `gorm:"not null" json:"first_name"`
	FamilyName string `gorm:"not null" json:"family_name"`
	BirthDate  datatypes.Date `gorm:"not null" json:"birth_date"`
	DeathDate  datatypes.Date `json:"death_date"`
	LifeSpan   string `json:"life_span"`
  Books      []Book `json:"books"`
}

type Genre struct {
  ID   uint `json:"id"`
	Name string `gorm:"unique_index;not null" json:"name"`
}

// `Book` belongs to `Author`, `AuthorID` is the foreign key
// `Book` belongs to `Genre`, `GenreID` is the foreign key
type Book struct {
	ID        uint `json:"id"`
	Title     string `gorm:"unique_index;not null" json:"title"`
  AuthorId  uint  `json:"author_id"`
	Author    Author `json:"author"` // `gorm:"foreignkey:AuthorID"`
	Summary   string `gorm:"not null" json:"summary"`
	ISBN      string `gorm:"not null" json:"isbn"`
  GenreId   uint   `json:"genre_id"`
	Genre     Genre `json:"genre"`  // `gorm:"foreignkey:ID"`
}

// `BookInstance` belongs to `Book`, `BookID` is the foreign key
type BookInstance struct {
	ID      uint `json:"id"`
  BookId  uint
	Book    Book `json:"book"` // `gorm:"foreignkey:ID"`
	Imprint string `gorm:"not null" json:"imprint"`
	Status  BookInstanceStatus `gorm:"not null" json:"status"`
  DueBack datatypes.Date `gorm:"not null" json:"due_date"`
}

type BookInstanceStatus uint8
const (
	NotAvailable BookInstanceStatus = iota
	OnOrder
	InTransit
	OnHold
	OnLoan
	InLibrary
)
