package store

import (
	"testing"
	"time"

	"github.com.vcholak.library/book"
	"github.com.vcholak.library/db"

	"gorm.io/datatypes"
	. "gorm.io/gorm/utils/tests"
)

var d = db.TestDB()

func TestCreateAuthor2(t *testing.T) {

	d.Migrator().DropTable(&book.Author{})
	if err := d.Migrator().AutoMigrate(&book.Author{}); err != nil {
		t.Errorf("failed to migrate, got error: %v", err)
	}

	createAuthor()

	result := book.Author{}
	if err := d.First(&result, "first_name = ? AND family_name = ?", "Xxx", "Yyy").Error; err != nil {
		t.Fatalf("Failed to find author")
	}

	AssertEqual(t, result.ID, 1)
}

func TestFindAll(t *testing.T) {
	d.Migrator().DropTable(&book.Author{})
	if err := d.Migrator().AutoMigrate(&book.Author{}); err != nil {
		t.Errorf("failed to migrate, got error: %v", err)
	}

	createAuthor()

	result := []book.Author{}
	if err := d.Find(&result).Error; err != nil {
		t.Fatalf("Failed to find all authors")
	}

	AssertEqual(t, len(result), 1)
}

func createAuthor() {
	author := book.Author{
		FirstName:  "Xxx",
		FamilyName: "Yyy",
		BirthDate:  datatypes.Date(time.Now().UTC()),
	}

	d.Create(&author)
}
