package repository

import "EuFeeding/entity"

type AnimalRepository interface {
	Add(a entity.Animal) error
	List(chatID int64) ([]entity.Animal, error)
}

type AnimalRepo struct {
}

func NewAnimalRepo() *AnimalRepo {
	return &AnimalRepo{}
}

func (ar *AnimalRepo) Add(a entity.Animal) error {
	//TODO implement me
	return nil
}

func (ar *AnimalRepo) List(chatID int64) ([]entity.Animal, error) {
	//TODO implement me
	return []entity.Animal{}, nil
}
