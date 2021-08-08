package store

import (
	"testing"

	"github.com.vcholak.library/model"
	. "gorm.io/gorm/utils/tests"
)

func TestNewGenre(t *testing.T) {
  migrateGenre(t)

  genre := createGenre("Test")
  AssertEqual(t, genre.ID, 1)
}

func TestFindGenreByName(t *testing.T) {
  migrateGenre(t)

  createGenre("Other")
  result := model.Genre{}
  if err := d.First(&result, "Name = ?", "Other").Error; err != nil {
    t.Fatalf("Failed to find genre")
  }
  AssertEqual(t, result.ID, 1)
}


func migrateGenre(t *testing.T) {
  d.Migrator().DropTable(&model.Genre{})
	if err := d.Migrator().AutoMigrate(&model.Genre{}); err != nil {
		t.Errorf("Failed to migrate, got error: %v", err)
	}
}

func createGenre(name string) model.Genre {
  genre := model.Genre{
    Name: name,
  }
  d.Create(&genre)
  return genre
}
