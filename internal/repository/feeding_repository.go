package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	"github.com/qrave1/PetFeedingBot/internal/repository/dto/v1"
)

type FeedingRepository interface {
	Create(ctx context.Context, feeding entity.Feeding) error
	GetForPet(ctx context.Context, petID string, month time.Time) ([]entity.Feeding, error)
	GetForChat(ctx context.Context, chatID int64, month time.Time) ([]entity.Feeding, error)
	Delete(ctx context.Context, id string) error
}

type feedingRepository struct {
	db *sqlx.DB
}

func NewFeedingRepository(db *sqlx.DB) FeedingRepository {
	return &feedingRepository{db: db}
}

func (r *feedingRepository) Create(ctx context.Context, feeding entity.Feeding) error {
	feedingDTO := v1.Feeding{
		ID:        feeding.ID,
		PetID:     feeding.PetID,
		FeedingAt: feeding.FeedingAt,
		FoodType:  feeding.FoodType,
	}

	query := `
		INSERT INTO feeding (id, pet_id, feeding_at, food_type, created_at)
		VALUES (:id, :pet_id, :feeding_at, :food_type, :created_at)`

	_, err := r.db.NamedExecContext(ctx, query, feedingDTO)

	return err
}

func (r *feedingRepository) GetForPet(ctx context.Context, petID string, month time.Time) ([]entity.Feeding, error) {
	var feedingDTO []v1.Feeding

	query := `SELECT * FROM feeding 
         WHERE pet_id = $1 
           AND strftime('%m', feeding_at) =$2 
           AND strftime('%Y', feeding_at) = $3;`

	err := r.db.SelectContext(ctx, &feedingDTO, query, petID, month.Month(), month.Year())

	if err != nil {
		return nil, err
	}

	feedings := make([]entity.Feeding, 0, len(feedingDTO))

	for _, feeding := range feedingDTO {
		feedings = append(feedings,
			entity.Feeding{
				ID:        feeding.ID,
				PetID:     feeding.PetID,
				FeedingAt: feeding.FeedingAt,
				FoodType:  feeding.FoodType,
			},
		)
	}

	return feedings, nil
}

func (r *feedingRepository) GetForChat(ctx context.Context, chatID int64, month time.Time) ([]entity.Feeding, error) {
	var feedingDTO []v1.Feeding

	query := `
	SELECT 
		feeding.id,
		feeding.pet_id,
		feeding.feeding_at,
		feeding.food_type
	FROM feeding
         JOIN pets on pets.id = feeding.pet_id
	WHERE pets.chat_id = $1
  	AND CAST(strftime('%m', feeding.feeding_at / 1000, 'unixepoch') AS INTEGER) = $2
  	AND CAST(strftime('%Y', feeding.feeding_at / 1000, 'unixepoch') AS INTEGER) = $3;`

	err := r.db.SelectContext(ctx, &feedingDTO, query, chatID, month.Month(), month.Year())

	if err != nil {
		return nil, err
	}

	feedings := make([]entity.Feeding, 0, len(feedingDTO))

	for _, feeding := range feedingDTO {
		feedings = append(feedings,
			entity.Feeding{
				ID:        feeding.ID,
				PetID:     feeding.PetID,
				FeedingAt: feeding.FeedingAt,
				FoodType:  feeding.FoodType,
			},
		)
	}

	return feedings, nil
}

func (r *feedingRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM feeding WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
