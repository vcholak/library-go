package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	authors := v1.Group("/authors")
	authors.HEAD("", h.AuthorsTotal)
	authors.GET("", h.Authors)
	authors.POST("", h.CreateAuthor)

	genres := v1.Group("/genres")
	genres.HEAD("", h.GenresTotal)
	genres.GET("", h.Genres)

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
}
