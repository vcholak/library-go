package db

import (
	"fmt"

	"os"

	"github.com.vcholak.library/model"
	"github.com.vcholak.library/utils"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func New() *gorm.DB {
  dataDir := utils.EnvVar("DATA_DIR", false)
  connetion_str := fmt.Sprintf("%s/library.db?cache=shared", dataDir)
	db, err := gorm.Open(sqlite.Open(connetion_str), &gorm.Config{
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
  dataDir := utils.EnvVar("DATA_DIR", true)
  connetion_str := fmt.Sprintf("%s/library_test.db", dataDir)
	db, err := gorm.Open(sqlite.Open(connetion_str), &gorm.Config{
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
  dataDir := utils.EnvVar("DATA_DIR", true)
  connetion_str := fmt.Sprintf("%s/library_test.db", dataDir)
	if err := os.Remove(connetion_str); err != nil {
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
