package v1

import "time"

type Feeding struct {
	ID        string    `db:"id"`
	PetID     string    `db:"pet_id"`
	FeedingAt time.Time `db:"feeding_at"`
	FoodType  string    `db:"food_type"`
}
