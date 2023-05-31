// DO NOT GENERATED FILE

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	ID        int
	Name      string
	Overview  string
	Timestamp time.Time
}

type BookHandler struct {
	tmpl *template.Template
	repo *Repo
}

func NewBookHandler(repo *Repo, tmpl *template.Template) (h *BookHandler, err error) {
	createSQL := `CREATE TABLE IF NOT EXISTS Books ( 
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT, 
                overview TEXT, 
                timestamp DATETIME NOT NULL
                )`

	if _, err := repo.Exec(createSQL); err != nil {
		return nil, err
	}

	h = &BookHandler{
		tmpl: tmpl,
		repo: repo,
	}

	return h, nil
}

func (h *BookHandler) RegisterHandlers(rtr *mux.Router) {
	rtr.HandleFunc("/books", h.getBooks).Methods("GET")
	rtr.HandleFunc("/books/{id}", h.getBook).Methods("GET")
	rtr.HandleFunc("/books/{id}", h.submitBook).Methods("POST")
}

func (h *BookHandler) getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.repo.Query("SELECT * FROM Books ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		book := Book{}

		err = rows.Scan(&book.ID, &book.Name, &book.Overview, &book.Timestamp)
		if err != nil {
			panic(err)
		}

		books = append(books, book)
	}

	h.tmpl.ExecuteTemplate(w, "books.html", books)
}

func (h *BookHandler) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id == 0 {
		h.tmpl.ExecuteTemplate(w, "book.html", &Book{})
	} else {
		row := h.repo.QueryRow("SELECT * FROM Books WHERE (id = ?)", id)

		book := &Book{}

		err := row.Scan(&book.ID, &book.Name, &book.Overview, &book.Timestamp)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Book not Found")
		} else {
			h.tmpl.ExecuteTemplate(w, "book.html", book)
		}
	}
}

func (h *BookHandler) submitBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	name := r.Form.Get("name")
	overview := r.Form.Get("overview")

	if id == 0 {
		execSQL := "INSERT INTO Books VALUES (NULL, ?, ?, ?);"
		_, err := h.repo.Exec(execSQL, name, overview, time.Now())
		if err != nil {
			panic(err)
		}
	} else {
		execSQL := "UPDATE Books SET name = ?, overview = ?, timestamp = ? WHERE (id = ?);"
		_, err := h.repo.Exec(execSQL, name, overview, time.Now(), id)
		if err != nil {
			panic(err)
		}
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
