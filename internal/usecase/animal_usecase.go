package usecase

import (
	"EuFeeding/internal/domain/entity"
	"EuFeeding/internal/domain/errs"
	"EuFeeding/internal/repository"
)

type AnimalUsecase interface {
	Add(chatID int64, name string) error
	List(chatID int64) ([]entity.Animal, error)
}

type AnimalUsecaseImpl struct {
	animalRepo repository.AnimalRepository
}

func NewAnimalUsecaseImpl(animalRepo repository.AnimalRepository) *AnimalUsecaseImpl {
	return &AnimalUsecaseImpl{
		animalRepo: animalRepo,
	}
}

func (a *AnimalUsecaseImpl) Add(chatID int64, name string) error {
	animal := entity.Animal{
		ChatID: chatID,
		Name:   name,
	}

	return a.animalRepo.Add(animal)
}

func (a *AnimalUsecaseImpl) List(chatID int64) ([]entity.Animal, error) {
	list, err := a.animalRepo.List(chatID)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errs.AnimalsNotFound
	}

	return list, nil
}
