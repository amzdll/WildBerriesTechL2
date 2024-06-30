package http_utils

import (
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, code int) {
	http.Error(
		w,
		fmt.Sprintf("{err: %s}", http.StatusText(code)),
		code,
	)
}
