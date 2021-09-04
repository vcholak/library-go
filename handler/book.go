package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com.vcholak.library/model"
	"github.com/labstack/echo/v4"
)

// BooksTotal returns total number of Books
func (h *Handler) BooksTotal(c echo.Context) error {

	books, err := h.bookStore.BookCount()
  if err != nil {
    fmt.Println("BooksTotal error:", err)
    return c.JSON(http.StatusInternalServerError, err)
  }

	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(books, 10))
  c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Books returns a list of all books
func (h *Handler) Books(c echo.Context) error {

  books, err := h.bookStore.Books()
  if err != nil {
    fmt.Println("Books error:", err)
    return c.JSON(http.StatusInternalServerError, err)
  }

	return c.JSON(http.StatusOK, books)
}

func (h *Handler) GetBook(c echo.Context) error {
	s := c.Param("id")

  id, err := strconv.ParseUint(s, 10, 64)
  if err != nil {
    fmt.Println("GetBook error:", err)
    return c.JSON(http.StatusBadRequest, err)
  }

  book, err2 := h.bookStore.GetBook(id)
  if err2 != nil {
    fmt.Println("GetBook error:", err2)
    return c.JSON(http.StatusInternalServerError, err2)
  }

  return c.JSON(http.StatusOK, book)
}

func (h *Handler) CreateBook(c echo.Context) error {
	book := new(model.Book)
  if err := c.Bind(book); err != nil {
    fmt.Println("CreateBook error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

  author, err := h.authorStore.GetAuthor(uint64(book.AuthorId))
  if err != nil {
    fmt.Println("CreateBook error:", err)
		return c.JSON(http.StatusInternalServerError, err)
  }
  book.Author = author

  genre, err := h.genreStore.GetGerne(uint64(book.GenreId))
  if err != nil {
    fmt.Println("CreateBook error:", err)
		return c.JSON(http.StatusInternalServerError, err)
  }
  book.Genre = genre

  if err := h.bookStore.NewBook(book); err != nil {
    fmt.Println("CreateBook error:", err)
		return c.JSON(http.StatusInternalServerError, err)
  }
  fmt.Println("Created book:", book)
	return c.JSON(http.StatusOK, book)
}

func (h *Handler) UpdateBook(c echo.Context) error {
	//
	return nil
}

func (h *Handler) DeleteBook(c echo.Context) error {
	//
	return nil
}
