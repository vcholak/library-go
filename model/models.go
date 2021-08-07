package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Author struct {
  ID uint `gorm:"primarykey" json:"id"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Firstname  string `gorm:"not null" json:"first_name"`
	Familyname string `gorm:"not null" json:"family_name"`
	Birth      datatypes.Date `gorm:"not null" json:"birth_date"`
	Death      datatypes.Date `json:"death_date"`
	Name       string `json:"name"`
	Lifespan   string `json:"life_span"`
}

type Genre struct {
	ID uint `gorm:"primarykey" json:"id"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Name string `gorm:"not null" json:"name"`
}

type Book struct {
	gorm.Model
	Title   string   `gorm:"unique_index;not null"`
	Authors []Author `gorm:"foreignkey:ID"`
	Summary string
	ISBN    string
	Genres  []Genre  `gorm:"many2many:book_genres;"`
}

type BookInstance struct {
	gorm.Model
	Book Book `gorm:"foreignkey:ID"`
	Imprint string
	Status  BookInstanceStatus `gorm:"not null"`
	Dueback string `sql:"type:date"`
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
