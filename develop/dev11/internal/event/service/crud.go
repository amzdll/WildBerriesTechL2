package service

import (
	"time"
)

func (s *Service) CreateDayEvent(e Event) error {
	if err := s.repository.CreateEvent(e); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetDayEvents(userId int, date time.Time) ([]Event, error) {
	events, err := s.repository.GetEvents(userId, date)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Service) GetWeekEvents(userId int, date time.Time) ([]Event, error) {
	day := date
	events := make([]Event, 0)

	for i := 0; i < 6; i++ {
		e, err := s.repository.GetEvents(userId, day)
		if err != nil {
			return nil, err
		}
		events = append(events, e...)
		day = day.Add(24 * time.Hour)
	}

	return events, nil
}

func (s *Service) GetMonthEvents(userId int, date time.Time) ([]Event, error) {
	day := date
	events := make([]Event, 0)

	for i := 0; i < 30; i++ {
		e, err := s.repository.GetEvents(userId, day)
		if err != nil {
			return nil, err
		}
		events = append(events, e...)
		day = day.Add(24 * time.Hour)
	}

	return events, nil
}

func (s *Service) UpdateDayEvent(e Event) error {
	return s.repository.UpdateEvent(e)
}

func (s *Service) DeleteDayEvent(e Event) error {
	return s.repository.DeleteEvent(e)
}
