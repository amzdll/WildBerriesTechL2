package handler

import (
	"net/http"
)

type Service interface {
	GetDayEvents()
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

	handler.HandleFunc("GET /events_for_day", h.getDayEvents)

	return handler
}
