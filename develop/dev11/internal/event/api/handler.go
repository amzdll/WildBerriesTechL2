package api

import (
	"l2/develop/dev11/internal/event/service"
	"net/http"
)

type Service interface {
	CreateDayEvent(event service.Event) error
	GetDayEvents()
	GetWeekEvents()
	GetMonthEvents()
	UpdateDayEvent()
	DeleteDayEvent()
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
