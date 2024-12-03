package db

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(rawDb *sql.DB) *Repository {
	return &Repository{db: rawDb}
}
