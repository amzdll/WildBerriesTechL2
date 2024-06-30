package db

import (
	"errors"
	"l2/develop/dev11/internal/event/service"
	"time"
)

type EventStorage map[int]UserStorage

type UserStorage map[time.Time][]service.Event

func (storage *EventStorage) Add(e service.Event) error {
	if _, ok := (*storage)[e.UserId]; !ok {
		(*storage)[e.UserId] = make(UserStorage)
		(*storage)[e.UserId][e.Date] = make([]service.Event, 0)
	}
	(*storage)[e.UserId][e.Date] = append((*storage)[e.UserId][e.Date], e)

	return nil
}

func (storage *EventStorage) Get(userId int, date time.Time) ([]service.Event, error) {
	if userStorage, ok := (*storage)[userId]; ok {
		if value, ok := userStorage[date]; ok {
			return value, nil
		}
	}

	return nil, nil
}

func (storage *EventStorage) Update(e service.Event) error {
	if userStorage, ok := (*storage)[e.UserId]; !ok {
		return errors.New("event not found")
	} else if events, ok := userStorage[e.Date]; !ok {
		return errors.New("event not found")
	} else {
		for i, event := range events {
			if event.Id == e.Id {
				(*storage)[e.UserId][e.Date][i] = e
				return nil
			}
		}
	}

	return errors.New("event not found")
}

func (storage *EventStorage) Delete(e service.Event) error {
	if userStorage, ok := (*storage)[e.UserId]; !ok {
		return errors.New("event not found")
	} else if events, ok := userStorage[e.Date]; !ok {
		return errors.New("event not found")
	} else {
		for i, event := range events {
			if event == e {
				(*storage)[e.UserId][e.Date] = append(events[:i], events[i+1:]...)
				return nil
			}
		}
	}

	return errors.New("event not found")
}
