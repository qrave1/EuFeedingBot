package usecase

import (
	"EuFeeding/internal/domain/entity"
	"EuFeeding/internal/domain/errs"
	"EuFeeding/internal/repository"
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
		ChatID: chatID,
		Name:   name,
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
