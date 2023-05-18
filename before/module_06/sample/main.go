package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Title     string
	Completed bool
}

type CoursePageData struct {
	PageTitle string
	Courses   []Course
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := CoursePageData{
			PageTitle: "My Courses",
			Courses: []Course{
				{Title: "Go kit", Completed: true},
				{Title: "Errors in Go", Completed: true},
				{Title: "Code Generation in Go", Completed: false},
			},
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
