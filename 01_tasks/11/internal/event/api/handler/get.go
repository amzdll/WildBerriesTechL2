package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) getDayEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s: Ивенты дня %s\n", r.URL.String(), r.Method)
}

func (h *Handler) getWeekEvents(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getMonthEvents(w http.ResponseWriter, r *http.Request) {

}
