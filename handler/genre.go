package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com.vcholak.library/model"
	"github.com/labstack/echo/v4"
)

// GenresTotal returns total number of Genres
func (h *Handler) GenresTotal(c echo.Context) error {

  genres, err := h.genreStore.GenreCount()
  if err != nil {
    fmt.Println("GenresTotal error:", err)
    return err
  }

	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(genres, 10))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Genres returns a list of all genres
func (h *Handler) Genres(c echo.Context) error {

  genres, err := h.genreStore.Genres()
  if err != nil {
    fmt.Println("Genres error:", err)
    return err
  }

	return c.JSON(http.StatusOK, genres)
}

// CreateGenre creates a new gerne and returns its ID
func (h *Handler) CreateGenre(c echo.Context) error {

  genre := new(model.Genre)
  if err := c.Bind(genre); err != nil {
    fmt.Println("CreateGenre error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}
  h.genreStore.NewGenre(genre)
  fmt.Println("Created gerne:", genre)
	return c.JSON(http.StatusOK, genre)
}

func (h *Handler) GetGenre(c echo.Context) error {

  s := c.Param("id")

  id, err := strconv.ParseUint(s, 10, 64)
  if err != nil {
    msg := fmt.Sprintf("Failed to parse id param: %s", s)
    fmt.Println("GetGenre error:", msg)
    return c.String(http.StatusBadRequest, msg)
  }

  genre, err2 := h.genreStore.GetGerne(id)
  if err2 != nil {
    msg := fmt.Sprintf("Failed to parse id param: %v", id)
    fmt.Println("GetGenre error:", msg)
    return err2
  }

  return c.JSON(http.StatusOK, genre)
}
