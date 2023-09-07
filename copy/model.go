package copy

import (
	"github.com.vcholak.library/book"

	"gorm.io/datatypes"
)

// `BookInstance` belongs to `Book`, `BookID` is the foreign key
type BookInstance struct {
	ID      uint `json:"id"`
	BookId  uint
	Book    book.Book          `json:"book"` // `gorm:"foreignkey:ID"`
	Imprint string             `gorm:"not null" json:"imprint"`
	Status  BookInstanceStatus `gorm:"not null" json:"status"`
	DueBack datatypes.Date     `gorm:"not null" json:"due_date"`
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
