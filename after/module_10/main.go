//go:generate go run gen/handler.go Book gen/book.txt gen/handler.tmpl book.go
//go:generate go run gen/handler.go Review gen/review.txt gen/handler.tmpl review.go

package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()

	repo, err := NewRepo("store.db")
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"inc": func(val int) int {
			return val + 1
		},
		"formatDate": func(timestamp time.Time) string {
			return timestamp.Format("02 Jan 2006")
		},
		"getColor": func(rating int) string {
			if rating == 1 {
				return "#ff6961"
			}
			if rating == 3 {
				return "#77dd77"
			}

			return "#aec6cf"
		},
		"calcDueDate": func(timestamp time.Time, rating int) time.Time {
			days := 14
			if rating == 1 {
				days = 7
			}
			if rating == 3 {
				days = 21
			}

			return timestamp.AddDate(0, 0, days)
		},
	}

	tmpl := template.Must(template.New("handler").Funcs(funcMap).ParseGlob("html/*"))

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
