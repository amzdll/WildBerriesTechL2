package event

import (
	"l2/01_tasks/11/internal/event/api/handler"
	"l2/01_tasks/11/internal/event/service"
)

func New() *handler.Handler {
	s := service.New()
	return handler.New(s)
}
