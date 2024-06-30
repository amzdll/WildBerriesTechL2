package repository

import "l2/develop/dev11/internal/db"

type Repository struct {
	db *db.EventStorage
}

func New(db *db.EventStorage) *Repository {
	return &Repository{db: db}
}
