package api

import (
	"l2/develop/dev11/internal/event/service"
	"net/http"
	"time"
)

type Service interface {
	CreateDayEvent(e service.Event) error
	GetDayEvents(userId int, date time.Time) ([]service.Event, error)
	GetWeekEvents(userId int, date time.Time) ([]service.Event, error)
	GetMonthEvents(userId int, date time.Time) ([]service.Event, error)
	UpdateDayEvent(e service.Event) error
	DeleteDayEvent(e service.Event) error
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("POST /create_event", h.createDayEvent)
	handler.HandleFunc("POST /update_event", h.updateDayEvent)
	handler.HandleFunc("POST /delete_event", h.deleteDayEvent)
	handler.HandleFunc("GET /events_for_day", h.getDayEvents)
	handler.HandleFunc("GET /events_for_week", h.getWeekEvents)
	handler.HandleFunc("GET /events_for_month", h.getMonthEvents)

	return handler
}
