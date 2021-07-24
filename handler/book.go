package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com.vcholak.library/model"
)

type BookResponse struct {
  ID uint `json:"id"`
  Title string `json:"title"`
}

// BooksTotal returns total number of Books
func (h *Handler) BooksTotal(c echo.Context) error {

	books := 10

	c.Response().Header().Set("X-Result-Count", strconv.Itoa(books))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Books returns a list of all books
func (h *Handler) Books(c echo.Context) error {

  //TODO call DB
  book := model.Book{
    Model: gorm.Model{ID: 1},
    Title: "The Talisman",
  }

  books := make([]BookResponse, 0)
  books = append(books, BookResponse{
    ID: book.ID,
    Title: book.Title,
  })

	return c.JSON(http.StatusOK, books)
}

func (h *Handler) GetBook(c echo.Context) error {
	//
	return nil
}

func (h *Handler) CreateBook(c echo.Context) error {
	//
	return nil
}

func (h *Handler) UpdateBook(c echo.Context) error {
	//
	return nil
}

func (h *Handler) DeleteBook(c echo.Context) error {
	//
	return nil
}
