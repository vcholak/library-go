package main

import (
	"github.com.vcholak.library/db"
	"github.com.vcholak.library/handler"
	"github.com.vcholak.library/router"
	"github.com.vcholak.library/store"
)

func main() {

	d := db.New()
	db.AutoMigrate(d)

	r := router.New()
	v1 := r.Group("/api")

	bs := store.NewBookStore(d)
	as := store.NewAuthorStore(d)
	cs := store.NewBookInstanceStore(d)
	gs := store.NewGerneStore(d)

	h := handler.NewHandler(bs, as, cs, gs)

	h.Register(v1)

	r.Logger.Fatal(r.Start(":8080"))
}
