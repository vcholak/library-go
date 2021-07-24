package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GenreResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

// GenresTotal returns total number of Genres
func (h *Handler) GenresTotal(c echo.Context) error {

  genres := 10

	c.Response().Header().Set("X-Result-Count", strconv.Itoa(genres))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Genres returns a list of all genres
func (h *Handler) Genres(c echo.Context) error {

  genre := GenreResponse{
    ID: 1,
    Name: "fiction",
    Url: "https://en.wikipedia.org/wiki/Fiction",
  }

	genres := make([]GenreResponse, 0)
  genres = append(genres, genre)

	return c.JSON(http.StatusOK, genres)
}
