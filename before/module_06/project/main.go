//go:generate go run gen/struct.go Book gen/book.txt gen/book.tmpl book.go

package main

import (
	"html/template"
	"net/http"
)

func main() {
	books := []Book{
		{ID: 1, Name: "The Alchemist"},
		{ID: 2, Name: "Atomic Habits"},
	}

	tmpl := template.Must(template.ParseFiles("book.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, books)
	})

	http.ListenAndServe(":8080", nil)
}
