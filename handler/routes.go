package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com.vcholak.library/book"
	"github.com.vcholak.library/copy"
	"github.com.vcholak.library/genre"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	authors := v1.Group("/authors")
	authors.HEAD("", h.AuthorsTotal)
	authors.GET("", h.Authors)
	authors.POST("", h.CreateAuthor)
	authors.GET("/:id", h.GetAuthor)

	genres := v1.Group("/genres")
	genres.HEAD("", h.GenresTotal)
	genres.GET("", h.Genres)
	genres.GET("/:id", h.GetGenre)
	genres.POST("", h.CreateGenre)

	books := v1.Group("/books")
	books.HEAD("", h.BooksTotal)
	books.GET("", h.Books)
	books.GET("/:id", h.GetBook)
	books.POST("", h.CreateBook)
	books.PUT("/:id", h.UpdateBook)
	books.DELETE("/:id", h.DeleteBook)

	copies := v1.Group("/copies")
	copies.HEAD("", h.CopiesTotal)
	copies.HEAD("/available", h.AvailableCopiesTotal)
	copies.GET("", h.Copies)
	copies.POST("", h.CreateBookCopy)
}

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

// GetBook returns a book
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

// CreateBook creates a book
func (h *Handler) CreateBook(c echo.Context) error {
	book := new(book.Book)
	if err := c.Bind(book); err != nil {
		fmt.Println("CreateBook error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	author, err := h.authorStore.GetAuthor(uint64(book.AuthorId))
	if err != nil {
		fmt.Println("GetAuthor error:", err)
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

// ------------------------
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
	author := new(book.Author)
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

// GetAuthor returns the author details
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

// ---------------------------------------
// CopiesTotal returns total number of BookInstances
func (h *Handler) CopiesTotal(c echo.Context) error {

	copies, err := h.copyStore.BookInstanceCount()
	if err != nil {
		panic("Failed to fetch book instance count")
	}

	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(copies, 10))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// AvailableCopiesTotal returns total number of BookInstances
func (h *Handler) AvailableCopiesTotal(c echo.Context) error {

	available_copies, err := h.copyStore.AvailableBookInstanceCount()
	if err != nil {
		panic("Failed to fetch book instance count")
	}

	c.Response().Header().Set("X-Result-Count", strconv.FormatInt(available_copies, 10))
	c.Response().Header().Set("Access-Control-Expose-Headers", "X-Result-Count")

	return c.NoContent(http.StatusOK)
}

// Copies returns all BookInstances
func (h *Handler) Copies(c echo.Context) error {
	copies, err := h.copyStore.BookInstances()
	if err != nil {
		fmt.Println("Copies error:", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, copies)
}

// CreateBookCopy creates a new BookInstance
func (h *Handler) CreateBookCopy(c echo.Context) error {
	book_copy := new(copy.BookInstance)
	if err := c.Bind(book_copy); err != nil {
		fmt.Println("CreateBookCopy error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	book, err := h.bookStore.GetBook(uint64(book_copy.BookId))
	if err != nil {
		fmt.Println("GetBook error:", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	book_copy.Book = book

	if err := h.copyStore.NewBookInstance(book_copy); err != nil {
		fmt.Println("CreateBookCopy error:", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("Created book copy:", book_copy)
	return c.JSON(http.StatusOK, book_copy)
}

// -------------------------------------------------
// GenresTotal returns total number of Genres
func (h *Handler) GenresTotal(c echo.Context) error {

	genres, err := h.genreStore.GenreCount()
	if err != nil {
		fmt.Println("GenresTotal error:", err)
		return c.JSON(http.StatusInternalServerError, err)
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
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, genres)
}

// CreateGenre creates a new gerne and returns its ID
func (h *Handler) CreateGenre(c echo.Context) error {

	g := new(genre.Genre)
	if err := c.Bind(g); err != nil {
		fmt.Println("CreateGenre error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.genreStore.NewGenre(g); err != nil {
		fmt.Println("CreateGenre error:", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("Created gerne:", g)
	return c.JSON(http.StatusOK, g)
}

func (h *Handler) GetGenre(c echo.Context) error {

	s := c.Param("id")

	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println("GetGenre error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	genre, err2 := h.genreStore.GetGerne(id)
	if err2 != nil {
		fmt.Println("GetGenre error:", err2)
		return c.JSON(http.StatusInternalServerError, err2)
	}

	return c.JSON(http.StatusOK, genre)
}
