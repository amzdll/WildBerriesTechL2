package app

import "l2/develop/dev11/internal/db"

func (a *App) initDB() {
	a.db = db.New()
}
