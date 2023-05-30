package repository

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func New(DB *sql.DB) *Repository {
	return &Repository{DB: DB}
}
