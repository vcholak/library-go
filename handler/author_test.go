package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com.vcholak.library/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {

  birth_date := time.Now().UTC()
  str := birth_date.Format(time.RFC3339)

  var userJSON = `{"first_name":"First","family_name":"Last", "birth_date":"` + str + `"}`

  url := "/api/authors"

  req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(userJSON))
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
      t.Errorf("The request could not be created because of: %v", err)
  }
  rec := httptest.NewRecorder()
  e := echo.New()
  c := e.NewContext(req, rec)

  author := new(model.Author)
  c.JSON(http.StatusOK, author)

  res := rec.Result()
  defer res.Body.Close()

  if assert.NoError(t, c.Bind(author)) {
    assert.Equal(t, http.StatusOK, rec.Code)
  }
}
