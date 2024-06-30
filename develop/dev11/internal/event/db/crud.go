package db

import (
	"l2/develop/dev11/internal/event/service"
	"time"
)

func (r *Repository) CreateEvent(e service.Event) error {
	return r.db.Add(e)
}

func (r *Repository) GetEvents(userId int, date time.Time) ([]service.Event, error) {
	return r.db.Get(userId, date)
}

func (r *Repository) UpdateEvent(e service.Event) error {
	return r.db.Update(e)
}

func (r *Repository) DeleteEvent(e service.Event) error {
	return r.db.Delete(e)
}
