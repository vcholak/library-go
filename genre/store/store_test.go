package store

import (
	"testing"

	"github.com.vcholak.library/db"
	"github.com.vcholak.library/genre"

	. "gorm.io/gorm/utils/tests"
)

var d = db.TestDB()

func TestNewGenre(t *testing.T) {
	migrateGenre(t)

	genre := createGenre("Test")
	AssertEqual(t, genre.ID, 1)
}

func TestFindGenreByName(t *testing.T) {
	migrateGenre(t)

	createGenre("Other")
	result := genre.Genre{}
	if err := d.First(&result, "Name = ?", "Other").Error; err != nil {
		t.Fatalf("Failed to find genre")
	}
	AssertEqual(t, result.ID, 1)
}

func migrateGenre(t *testing.T) {
	d.Migrator().DropTable(&genre.Genre{})
	if err := d.Migrator().AutoMigrate(&genre.Genre{}); err != nil {
		t.Errorf("Failed to migrate, got error: %v", err)
	}
}

func createGenre(name string) genre.Genre {
	genre := genre.Genre{
		Name: name,
	}
	d.Create(&genre)
	return genre
}
