package db

import "l2/develop/dev11/internal/event/service"

type Event struct {
	Title string
	Body  string
}

func FromModel(e service.Event) Event {
	return Event{
		Title: e.Title,
		Body:  e.Body,
	}
}
