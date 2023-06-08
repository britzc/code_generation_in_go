// DO NOT EDIT GENERATED FILE

package main

import (
	"fmt"
	"html/template"

	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Review struct {
	ID        int
	BookID    int
	Content   string
	Rating    int
	Timestamp time.Time
}

type ReviewHandler struct {
	tmpl *template.Template
	repo *Repo
}

func NewReviewHandler(repo *Repo, tmpl *template.Template) (h *ReviewHandler, err error) {
	createSQL := `CREATE TABLE IF NOT EXISTS Reviews (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                bookid INTEGER,
                content TEXT,
                rating INTEGER,
                timestamp DATETIME NOT NULL
                )`

	if _, err := repo.Exec(createSQL); err != nil {
		return nil, err
	}

	h = &ReviewHandler{
		tmpl: tmpl,
		repo: repo,
	}

	return h, nil
}

func (h *ReviewHandler) RegisterHandlers(rtr *mux.Router) {
	rtr.HandleFunc("/reviews", h.getReviews).Methods("GET")
	rtr.HandleFunc("/reviews/{id}", h.getReview).Methods("GET")
	rtr.HandleFunc("/reviews/{id}", h.submitReview).Methods("POST")
}

func (h *ReviewHandler) getReviews(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	bookID, _ := strconv.Atoi(vals["bookid"][0])

	rows, err := h.repo.Query("SELECT * FROM Reviews WHERE (bookid = ?) ORDER BY ID ASC", bookID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	reviews := []Review{}
	for rows.Next() {
		review := Review{}

		err = rows.Scan(&review.ID, &review.BookID, &review.Content, &review.Rating, &review.Timestamp)
		if err != nil {
			panic(err)
		}

		reviews = append(reviews, review)
	}

	h.tmpl.ExecuteTemplate(w, "reviews.html", reviews)
}

func (h *ReviewHandler) getReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	vals := r.URL.Query()
	bookID, _ := strconv.Atoi(vals["bookid"][0])

	if id == 0 {
		review := &Review{}

		review.BookID = bookID
		h.tmpl.ExecuteTemplate(w, "review.html", &review)
	} else {
		row := h.repo.QueryRow("SELECT * FROM Reviews WHERE (id = ?)", id)

		review := &Review{}

		err := row.Scan(&review.ID, &review.BookID, &review.Content, &review.Rating, &review.Timestamp)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Review not Found")
		} else {

			h.tmpl.ExecuteTemplate(w, "review.html", review)
		}
	}
}

func (h *ReviewHandler) submitReview(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	bookid, _ := strconv.Atoi(r.Form.Get("bookid"))
	content := r.Form.Get("content")
	rating, _ := strconv.Atoi(r.Form.Get("rating"))

	if id == 0 {

		execSQL := "INSERT INTO Reviews VALUES (NULL, ?, ?, ?, ?);"

		_, err := h.repo.Exec(execSQL, bookid, content, rating, time.Now())
		if err != nil {
			panic(err)
		}

	} else {

		execSQL := "UPDATE Reviews SET bookid = ?, content = ?, rating = ?, timestamp = ? WHERE (id = ?);"

		_, err := h.repo.Exec(execSQL, bookid, content, rating, time.Now(), id)
		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, fmt.Sprintf("/reviews?bookid=%d", bookid), http.StatusSeeOther)

}
