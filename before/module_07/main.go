//go:generate go run gen/handler.go Book gen/book.txt gen/book.tmpl book.go

package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	tmpl := template.Must(template.ParseGlob("html/*"))

	bookHandler, err := NewBookHandler(tmpl)
	if err != nil {
		panic(err)
	}

	bookHandler.RegisterHandlers(rtr)

	http.ListenAndServe(":8080", rtr)
}
