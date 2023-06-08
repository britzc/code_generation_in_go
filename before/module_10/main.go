//go:generate go run gen/handler.go Book gen/book.txt gen/handler.tmpl book.go
//go:generate go run gen/handler.go Review gen/review.txt gen/handler.tmpl review.go

package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	repo, err := NewRepo("store.db")
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("html/*"))

	bookHandler, err := NewBookHandler(repo, tmpl)
	if err != nil {
		panic(err)
	}

	bookHandler.RegisterHandlers(rtr)

	reviewHandler, err := NewReviewHandler(repo, tmpl)
	if err != nil {
		panic(err)
	}

	reviewHandler.RegisterHandlers(rtr)

	http.ListenAndServe(":8080", rtr)
}
