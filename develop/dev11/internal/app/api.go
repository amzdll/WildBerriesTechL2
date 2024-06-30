package app

import (
	"l2/develop/dev11/internal/event"
	"l2/develop/dev11/internal/middleware"
	"net/http"
)

type Handler interface {
	Routes() *http.ServeMux
}

func (a *App) initAPI() {
	mux := http.NewServeMux()
	handlers := map[string]Handler{
		"/": event.New(),
	}

	for path, handler := range handlers {
		mux.Handle(path, handler.Routes())
	}

	loggingMiddleware := func(next http.Handler) http.Handler {
		return middleware.Logger(next, a.logger)
	}

	a.server = &http.Server{
		Addr: ":8000",
		Handler: buildMiddlewareChain(
			mux,
			loggingMiddleware,
		),
	}
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
