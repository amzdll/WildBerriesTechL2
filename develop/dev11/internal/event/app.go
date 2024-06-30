package event

import (
	"l2/develop/dev11/internal/event/api"
	"l2/develop/dev11/internal/event/db"
	"l2/develop/dev11/internal/event/service"
)

func New() *api.Handler {
	r := db.New()
	s := service.New(r)
	return api.New(s)
}
