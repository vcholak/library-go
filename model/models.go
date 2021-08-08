package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Author struct {
  gorm.Model
	Firstname  string `gorm:"not null" json:"first_name"`
	Familyname string `gorm:"not null" json:"family_name"`
	Birth      datatypes.Date `gorm:"not null" json:"birth_date"`
	Death      datatypes.Date `json:"death_date"`
	Name       string `json:"name"`
	Lifespan   string `json:"life_span"`
}

type Genre struct {
  gorm.Model
	Name string `gorm:"unique_index;not null" json:"name"`
}

// `Book` belongs to `Author`, `AuthorID` is the foreign key
// `Book` belongs to `Genre`, `GenreID` is the foreign key
type Book struct {
	gorm.Model
	Title   string `gorm:"unique_index;not null" json:"title"`
  AuthorID uint
	Author  Author // `gorm:"foreignkey:AuthorID"`
	Summary string `gorm:"not null" json:"summary"`
	ISBN    string `gorm:"not null" json:"isbn"`
  GenreID uint
	Genre   Genre  // `gorm:"foreignkey:ID"`
}

// `BookInstance` belongs to `Book`, `BookID` is the foreign key
type BookInstance struct {
	gorm.Model
  BookID uint
	Book Book // `gorm:"foreignkey:ID"`
	Imprint string `gorm:"not null" json:"inprint"`
	Status  BookInstanceStatus `gorm:"not null" json:"status"`
	Dueback string `sql:"type:date" gorm:"not null" json:"due_date"`
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
