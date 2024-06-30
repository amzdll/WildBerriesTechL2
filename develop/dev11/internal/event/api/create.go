package api

import (
	"encoding/json"
	"l2/develop/dev11/pkg/http_utils"
	"net/http"
	"time"
)

func (h *Handler) createDayEvent(w http.ResponseWriter, r *http.Request) {
	var e Event

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil || e.validate() != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-1-2", e.Date); err != nil {
		http_utils.Error(w, http.StatusBadRequest)
		return
	}

	if err := h.service.CreateDayEvent(e.toModel()); err != nil {
		http_utils.Error(w, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		return
	}
}
