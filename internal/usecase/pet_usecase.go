package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	"github.com/qrave1/PetFeedingBot/internal/domain/errs"
	"github.com/qrave1/PetFeedingBot/internal/repository"
)

type PetUsecase interface {
	Add(chatID int64, name string) error
	List(chatID int64) ([]entity.Pet, error)
}

type PetUsecaseImpl struct {
	petRepo repository.PetRepository
}

func NewPetUsecaseImpl(petRepo repository.PetRepository) *PetUsecaseImpl {
	return &PetUsecaseImpl{
		petRepo: petRepo,
	}
}

func (a *PetUsecaseImpl) Add(chatID int64, name string) error {
	pet := entity.Pet{
		ID:        uuid.New().String(),
		ChatID:    chatID,
		Name:      name,
		CreatedAt: time.Now(),
	}

	return a.petRepo.Add(pet)
}

func (a *PetUsecaseImpl) List(chatID int64) ([]entity.Pet, error) {
	list, err := a.petRepo.List(chatID)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errs.PetsNotFound
	}

	return list, nil
}
