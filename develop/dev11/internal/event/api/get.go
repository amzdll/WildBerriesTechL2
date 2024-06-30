package api

import (
	"encoding/json"
	"l2/develop/dev11/pkg/http_utils"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getDayEvents(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-1-2", r.URL.Query().Get("date"))
	if err != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	events, err := h.service.GetDayEvents(userId, date)
	if err != nil {
		http_utils.Error(w, http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(events); err != nil {
		return
	}
}

func (h *Handler) getWeekEvents(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-1-2", r.URL.Query().Get("date"))
	if err != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	events, err := h.service.GetWeekEvents(userId, date)
	if err != nil {
		http_utils.Error(w, http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(events); err != nil {
		return
	}
}

func (h *Handler) getMonthEvents(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-1-2", r.URL.Query().Get("date"))
	if err != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	events, err := h.service.GetWeekEvents(userId, date)
	if err != nil {
		http_utils.Error(w, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(events); err != nil {
		return
	}
}
