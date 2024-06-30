package app

import (
	"fmt"
	"l2/develop/dev11/internal/event/db"
	"net/http"
)

type App struct {
	logger *Logger
	server *http.Server
	db     *db.EventStorage
}

func New() *App {
	var app App

	app.initLogger()
	app.initAPI()

	return &app
}

func (a *App) Start() {
	a.logger.Info(fmt.Sprintf("Starting server on: %s", a.server.Addr))
	if err := a.server.ListenAndServe(); err != nil {
		a.logger.Fatal("err", err)
		return
	}
}
