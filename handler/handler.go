package handler

import (
	bookst "github.com.vcholak.library/book/store"
	copyst "github.com.vcholak.library/copy/store"
	genrest "github.com.vcholak.library/genre/store"
)

type Handler struct {
	bookStore   *bookst.BookStore
	authorStore *bookst.AuthorStore
	copyStore   *copyst.BookInstanceStore
	genreStore  *genrest.GenreStore
}

func NewHandler(bs *bookst.BookStore, as *bookst.AuthorStore, cs *copyst.BookInstanceStore,
    gs *genrest.GenreStore) *Handler {
	return &Handler{
		bookStore:   bs,
		authorStore: as,
		copyStore:   cs,
		genreStore:  gs,
	}
}
