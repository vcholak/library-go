package store

import (
	"testing"
	"time"

	"github.com.vcholak.library/db"
	"github.com.vcholak.library/model"

	"gorm.io/datatypes"
	. "gorm.io/gorm/utils/tests"
)

func TestCreate(t *testing.T)  {

  d := db.TestDB()

  d.Migrator().DropTable(&model.Author{})
	if err := d.Migrator().AutoMigrate(&model.Author{}); err != nil {
		t.Errorf("failed to migrate, got error: %v", err)
	}

  author := model.Author{
    ID: 1,
    Firstname: "Xxx",
    Familyname: "Yyy",
    Birth: datatypes.Date(time.Now().UTC()),
  }

  d.Create(&author)

  result := model.Author{}
	if err := d.First(&result, "firstName = ? AND familyName = ?", "Xxx", "Yyy").Error; err != nil {
		t.Fatalf("Failed to find record")
	}

	AssertEqual(t, result.ID, 1)
}
