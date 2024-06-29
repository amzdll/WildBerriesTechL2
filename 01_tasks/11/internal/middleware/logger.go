package middleware

import (
	"fmt"
	"net/http"
)

type Logger interface {
	Info(msg string)
}

func LoggingMiddleware(next http.Handler, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	})
}
