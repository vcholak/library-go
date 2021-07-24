package model

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Firstname  string `gorm:"not null"`
	Familyname string `gorm:"not null"`
	Birth      *time.Time `gorm:"not null"`
	Death      *time.Time 
	Name       string 
	Lifespan   string 
	Url        string
}

type Genre struct {
	gorm.Model
	Name string 
	Url  string
}

type Book struct {
	gorm.Model
	Title   string   `gorm:"unique_index;not null"`
	Authors []Author `gorm:"foreignkey:ID"`
	Summary string
	ISBN    string
	Genres  []Genre  `gorm:"many2many:book_genres;"`
	Url     string
}

type BookInstance struct {
	gorm.Model
	Book Book `gorm:"foreignkey:ID"`
	Imprint string
	Status  BookInstanceStatus `gorm:"not null"`
	Dueback string `sql:"type:date"`
	Url     string
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
