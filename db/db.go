package db

import (
	"fmt"

	"os"

	"github.com.vcholak.library/model"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
  if err != nil {
    fmt.Println("Storage error: ", err)
    panic("Failed to connect database")
  }
	db.Set("MaxIdleConns", 3)
	db.Set("LogMode", true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("library_test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
  if err != nil {
    fmt.Println("Storage error: ", err)
    panic("Failed to connect database")
  }
	db.Set("MaxIdleConns", 3)
	db.Set("LogMode", false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("library_test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Author{},
		&model.Genre{},
		&model.Book{},
		&model.BookInstance{},
	)
	if err != nil {
		fmt.Println("DB migration error: ", err)
    panic("Failed to do database migrations")
	}
}
