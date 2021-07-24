package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthorResponse struct {
	ID uint `json:"id"`
	Firstname  string `json:"first_name"`
	Familyname string `json:"family_name"`
	Birth      *time.Time `json:"birth_date"`
	Death      *time.Time `json:"death_date"`
	Name       string `json:"name"`
	Lifespan   string `json:"life_span"`
	Url        string `json:"url"`
}

// AuthorsTotal returns total number of Authors
func (h *Handler) AuthorsTotal(c echo.Context) error {

  copies := 10

	c.Response().Header().Set("X-Result-Count", strconv.Itoa(copies))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Authors returns a list of all authors
func (h *Handler) Authors(c echo.Context) error {

  birth_date, _ := time.Parse(time.ANSIC, "1771-Aug-15")
  death_date, _ := time.Parse(time.ANSIC, "1832-Sep-21")

  author := AuthorResponse{
    ID: 1,
    Firstname: "Walter",
    Familyname: "Scott",
    Birth: &birth_date,
    Death: &death_date,
    Name: "Walter Scott",
    Lifespan: "Sir Walter Scott was a Scottish historical novelist, poet, playwright, and historian.",
    Url: "https://en.wikipedia.org/wiki/Walter_Scott",
  }

  authors := make([]AuthorResponse, 0)
  authors = append(authors, author)

	return c.JSON(http.StatusOK, authors)
}

// CreateAuthor creates a new author
func (h *Handler) CreateAuthor(c echo.Context) error {
	//
	return nil
}
