package repository

import (
	"l2/develop/dev11/internal/db"
	"l2/develop/dev11/internal/event/service"
	"time"
)

func (r *Repository) CreateDayEvent(e service.Event) error {
	(*r.db)[e.UserId][time.Time{}] = append((*r.db)[e.UserId][time.Time{}], db.FromModel(e))
	return nil
}

func (r *Repository) GetDayEvents() {}

func (r *Repository) GetWeekEvents() {}

func (r *Repository) GetMonthEvents() {}

func (r *Repository) UpdateDayEvent() {}

func (r *Repository) DeleteDayEvent() {}
