package api

import "l2/develop/dev11/internal/event/service"

type Event struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (e *Event) toModel() service.Event {
	return service.Event{
		UserId: e.UserID,
		Title:  e.Title,
		Body:   e.Body,
	}
}
