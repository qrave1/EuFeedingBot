package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	"github.com/qrave1/PetFeedingBot/internal/repository"
)

type FeedingUsecase interface {
	Add(ctx context.Context, petID string, feedingAt time.Time, foodType string) error
	GetForPet(ctx context.Context, petID string) (entity.Feeding, error)
	Delete(ctx context.Context, id string) error
}

type FeedingUsecaseImpl struct {
	feedingRepo repository.FeedingRepository
}

func NewFeedingUsecaseImpl(feedingRepo repository.FeedingRepository) *FeedingUsecaseImpl {
	return &FeedingUsecaseImpl{feedingRepo: feedingRepo}
}

func (f *FeedingUsecaseImpl) Add(ctx context.Context, petID string, feedingAt time.Time, foodType string) error {
	feeding := entity.Feeding{
		ID:        uuid.New().String(),
		PetID:     petID,
		FeedingAt: feedingAt,
		FoodType:  foodType,
		CreatedAt: time.Now(),
	}

	err := f.feedingRepo.Create(ctx, feeding)
	if err != nil {
		return err
	}

	return nil
}

func (f *FeedingUsecaseImpl) GetForPet(ctx context.Context, petID string) (entity.Feeding, error) {
	return f.feedingRepo.GetForPet(ctx, petID)
}

func (f *FeedingUsecaseImpl) Delete(ctx context.Context, id string) error {
	return f.feedingRepo.Delete(ctx, id)
}
