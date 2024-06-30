package service

func (s *Service) CreateDayEvent(e Event) error {
	if err := s.repository.CreateDayEvent(e); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetDayEvents() {}

func (s *Service) GetWeekEvents() {}

func (s *Service) GetMonthEvents() {}

func (s *Service) UpdateDayEvent() {}

func (s *Service) DeleteDayEvent() {}
