package service

import "time"

type Event struct {
	Id     int
	UserId int
	Date   time.Time
	Title  string
	Body   string
}
