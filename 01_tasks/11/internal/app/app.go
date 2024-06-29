package app

import (
	"fmt"
	"l2/01_tasks/11/internal/logger"
	"net/http"
)

type Handler interface {
	Routes() *http.ServeMux
}

type App struct {
	logger *logger.Logger
	server *http.Server
}

func New() *App {
	l := logger.New()
	s := initServer(l)

	return &App{
		logger: l,
		server: s,
	}
}

func (a *App) Start() {
	a.logger.Info(fmt.Sprintf("Starting server on: %s", a.server.Addr))
	if err := a.server.ListenAndServe(); err != nil {
		a.logger.Fatal("err", err)
		return
	}
}
