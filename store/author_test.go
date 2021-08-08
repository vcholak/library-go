package store

import (
	"testing"
	"time"

	"github.com.vcholak.library/db"
	"github.com.vcholak.library/model"

	"gorm.io/datatypes"
	. "gorm.io/gorm/utils/tests"
)

var d = db.TestDB()

func TestCreateAuthor(t *testing.T) {

  d.Migrator().DropTable(&model.Author{})
	if err := d.Migrator().AutoMigrate(&model.Author{}); err != nil {
		t.Errorf("failed to migrate, got error: %v", err)
	}

  createAuthor()

  result := model.Author{}
	if err := d.First(&result, "Firstname = ? AND Familyname = ?", "Xxx", "Yyy").Error; err != nil {
		t.Fatalf("Failed to find author")
	}

	AssertEqual(t, result.ID, 1)
}

func TestFindAll(t *testing.T) {
  d.Migrator().DropTable(&model.Author{})
	if err := d.Migrator().AutoMigrate(&model.Author{}); err != nil {
		t.Errorf("failed to migrate, got error: %v", err)
	}

  createAuthor()

  result := []model.Author{}
  if err := d.Find(&result).Error; err != nil {
    t.Fatalf("Failed to find all authors")
  }

  AssertEqual(t, len(result), 1)
}

func createAuthor() {
  author := model.Author{
    Firstname: "Xxx",
    Familyname: "Yyy",
    Birth: datatypes.Date(time.Now().UTC()),
  }

  d.Create(&author)
}
