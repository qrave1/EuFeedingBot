package entity

import "time"

type Feeding struct {
	ID        string
	PetID     string
	FeedingAt time.Time
	FoodType  string
	CreatedAt time.Time
}
