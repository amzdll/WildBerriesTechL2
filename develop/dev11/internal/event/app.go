package event

import (
	"l2/develop/dev11/internal/db"
	"l2/develop/dev11/internal/event/api"
	"l2/develop/dev11/internal/event/repository"
	"l2/develop/dev11/internal/event/service"
)

func New(db *db.EventStorage) *api.Handler {
	r := repository.New(db)
	s := service.New(r)
	return api.New(s)
}
