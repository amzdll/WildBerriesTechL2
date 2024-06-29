package app

import (
	"l2/01_tasks/11/internal/event"
	"l2/01_tasks/11/internal/logger"
	"l2/01_tasks/11/internal/middleware"
	"net/http"
)

func initServer(logger *logger.Logger) *http.Server {
	mux := http.NewServeMux()
	handlers := map[string]Handler{
		"/": event.New(),
	}

	for path, handler := range handlers {
		mux.Handle(path, handler.Routes())
	}

	loggingMiddleware := func(next http.Handler) http.Handler {
		return middleware.LoggingMiddleware(next, logger)
	}

	server := &http.Server{
		Addr: ":8000",
		Handler: buildMiddlewareChain(
			mux,
			loggingMiddleware,
		),
	}

	return server
}

func buildMiddlewareChain(
	mux http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	for _, m := range middlewares {
		mux = m(mux)
	}
	return mux
}
