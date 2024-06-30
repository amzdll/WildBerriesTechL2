package service

import "time"

type Repository interface {
	CreateEvent(e Event) error
	GetEvents(userId int, date time.Time) ([]Event, error)
	UpdateEvent(e Event) error
	DeleteEvent(e Event) error
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
