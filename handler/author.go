package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com.vcholak.library/model"
	"github.com/labstack/echo/v4"
)

// AuthorsTotal returns total number of Authors
func (h *Handler) AuthorsTotal(c echo.Context) error {

  authors, err := h.authorStore.AuthorCount()
  if err != nil {
    fmt.Println("AuthorsTotal error:", err)
    return c.JSON(http.StatusInternalServerError, err)
  }

	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(authors, 10))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Authors returns a list of all authors
func (h *Handler) Authors(c echo.Context) error {

  authors, err := h.authorStore.Authors()
  if err != nil {
    fmt.Println("Authors error:", err)
    return c.JSON(http.StatusInternalServerError, err)
  }

	return c.JSON(http.StatusOK, authors)
}

// CreateAuthor creates a new author
func (h *Handler) CreateAuthor(c echo.Context) error {
	author := new(model.Author)
  if err := c.Bind(author); err != nil {
    fmt.Println("CreateAuthor error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}
  if err := h.authorStore.NewAuthor(author); err != nil {
    fmt.Println("CreateAuthor error:", err)
		return c.JSON(http.StatusInternalServerError, err)
  }
  fmt.Println("Created author:", author)
	return c.JSON(http.StatusOK, author)
}

// GetAuthor returns author details
func (h *Handler) GetAuthor(c echo.Context) error {

  s := c.Param("id")

  id, err := strconv.ParseUint(s, 10, 64)
  if err != nil {
    fmt.Println("GetAuthor error:", err)
    return c.JSON(http.StatusBadRequest, err)
  }

  author, err2 := h.authorStore.GetAuthor(id)
  if err2 != nil {
    fmt.Println("GetAuthor error:", err2)
    return c.JSON(http.StatusInternalServerError, err2)
  }

  return c.JSON(http.StatusOK, author)
}
