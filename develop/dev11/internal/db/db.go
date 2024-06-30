package db

import (
	"time"
)

type UserStorage map[time.Time][]Event

type EventStorage map[int]UserStorage

func New() *EventStorage {
	storage := make(EventStorage)
	return &storage
}
