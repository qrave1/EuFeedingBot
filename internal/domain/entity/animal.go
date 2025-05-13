package entity

import "time"

type Animal struct {
	ID        string
	ChatID    int64
	Name      string
	CreatedAt time.Time
}
