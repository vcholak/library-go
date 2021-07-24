package handler

import "github.com.vcholak.library/store"

type Handler struct {
	bookStore   *store.BookStore
	authorStore *store.AuthorStore
	copyStore   *store.BookInstanceStore
	genreStore  *store.GenreStore
}

func NewHandler(bs *store.BookStore, as *store.AuthorStore, cs *store.BookInstanceStore, gs *store.GenreStore) *Handler {
	return &Handler{
		bookStore:   bs,
		authorStore: as,
		copyStore:   cs,
		genreStore:  gs,
	}
}
