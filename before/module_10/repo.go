package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(file string) (repo *Repo, err error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	repo = &Repo{
		db: db,
	}

	return repo, nil
}

func (repo *Repo) Close() (err error) {
	return repo.db.Close()
}

func (repo *Repo) Exec(query string, args ...any) (res sql.Result, err error) {
	res, err = repo.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *Repo) Query(query string, args ...any) (rows *sql.Rows, err error) {
	rows, err = repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (repo *Repo) QueryRow(query string, args ...any) (row *sql.Row) {
	return repo.db.QueryRow(query, args...)
}
