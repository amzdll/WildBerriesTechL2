package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) createDayEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s: Ивенты дня %s\n", r.URL.String(), r.Method)
}
