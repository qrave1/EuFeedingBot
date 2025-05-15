package entity

import (
	"time"
)

type Pet struct {
	ID        string
	ChatID    int64
	Name      string
	CreatedAt time.Time
}
