package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	"github.com/qrave1/PetFeedingBot/internal/repository/dto/v1"
)

type FeedingRepository interface {
	Create(ctx context.Context, feeding entity.Feeding) error
	GetForPet(ctx context.Context, petID string) (entity.Feeding, error)
	//Update(ctx context.Context, feeding entity.Feeding) error
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
		CreatedAt: feeding.CreatedAt,
	}

	query := `
		INSERT INTO feeding (id, pet_id, feeding_at, food_type, created_at)
		VALUES (:id, :pet_id, :feeding_at, :food_type, :created_at)`

	_, err := r.db.NamedExecContext(ctx, query, feedingDTO)

	return err
}

func (r *feedingRepository) GetForPet(ctx context.Context, petID string) (entity.Feeding, error) {
	var feedingDTO v1.Feeding

	query := `SELECT * FROM feeding WHERE pet_id = $1`

	err := r.db.SelectContext(ctx, &feedingDTO, query, petID)

	if err != nil {
		return entity.Feeding{}, err
	}

	return entity.Feeding{
		ID:        feedingDTO.ID,
		PetID:     feedingDTO.PetID,
		FeedingAt: feedingDTO.FeedingAt,
		FoodType:  feedingDTO.FoodType,
		CreatedAt: feedingDTO.CreatedAt,
	}, nil
}

func (r *feedingRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM feeding WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
