package db

type Repository struct {
	db *EventStorage
}

func New() *Repository {
	db := make(EventStorage)
	return &Repository{db: &db}
}
