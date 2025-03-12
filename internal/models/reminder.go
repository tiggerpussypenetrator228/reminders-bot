package models

import "time"

type Reminder struct {
	ID          int
	UserID      int
	Content     string
	Interval    time.Duration
	IsActive    bool
	LastChecked time.Time
}
