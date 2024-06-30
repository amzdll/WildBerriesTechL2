package service

type Repository interface {
	CreateDayEvent(e Event) error
	GetDayEvents()
	GetWeekEvents()
	GetMonthEvents()
	UpdateDayEvent()
	DeleteDayEvent()
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
