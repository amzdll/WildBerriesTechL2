package api

import (
	"errors"
	"l2/develop/dev11/internal/event/service"
	"time"
)

type Event struct {
	Id     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (e *Event) toModel() service.Event {
	date, _ := time.Parse("2006-1-2", e.Date)
	return service.Event{
		Id:     e.Id,
		UserId: e.UserID,
		Date:   date,
		Title:  e.Title,
		Body:   e.Body,
	}
}

func (e *Event) validate() error {
	if e.Id == 0 {
		return errors.New("empty data")
	}
	if e.UserID == 0 {
		return errors.New("empty data")
	}
	if e.Date == "" {
		return errors.New("empty data")
	}
	if e.Title == "" {
		return errors.New("empty data")
	}
	if e.Body == "" {
		return errors.New("empty data")
	}
	return nil
}
