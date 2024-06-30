package api

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) createDayEvent(w http.ResponseWriter, r *http.Request) {
	var e Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err := h.service.CreateDayEvent(e.toModel())
	if err != nil {
		return
	}
	//h.Lock()
	//e.ID = h.idCounter
	//e.CreatedAt = time.Now()
	//e.UpdatedAt = time.Now()
	//h.idCounter++
	//h.events = append(h.events, e)
	//h.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		return
	}
}

func (h *Handler) getDayEvents(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getWeekEvents(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getMonthEvents(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) updateDayEvent(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) deleteDayEvent(w http.ResponseWriter, r *http.Request) {}
