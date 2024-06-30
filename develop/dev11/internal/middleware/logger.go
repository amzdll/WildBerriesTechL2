package middleware

import (
	"fmt"
	"net/http"
)

type Log interface {
	Info(msg string)
}

func Logger(next http.Handler, logger Log) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	})
}
