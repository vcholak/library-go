package main

import (
	bookst "github.com.vcholak.library/book/store"
	copyst "github.com.vcholak.library/copy/store"
	"github.com.vcholak.library/db"
	genrest "github.com.vcholak.library/genre/store"
	"github.com.vcholak.library/handler"
	"github.com.vcholak.library/router"
)

func main() {

	d := db.New()
	db.AutoMigrate(d)

	r := router.New()
	v1 := r.Group("/api")

	bs := bookst.NewBookStore(d)
	as := bookst.NewAuthorStore(d)
	cs := copyst.NewBookInstanceStore(d)
	gs := genrest.NewGerneStore(d)

	h := handler.NewHandler(bs, as, cs, gs)

	h.Register(v1)

	r.Logger.Fatal(r.Start(":8080"))
}
