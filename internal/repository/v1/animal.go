package v1

import "time"

type AnimalDTO struct {
	ID        string    `db:"id"`
	ChatID    int64     `db:"chat_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
