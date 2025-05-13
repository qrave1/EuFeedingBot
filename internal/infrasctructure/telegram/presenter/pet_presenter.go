package presenter

import (
	"fmt"

	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
)

type PetPresenter interface {
	ConvertPetsList(pets []entity.Pet) string
}

type petPresenter struct{}

func NewPetPresenter() PetPresenter {
	return &petPresenter{}
}

func (p *petPresenter) ConvertPetsList(pets []entity.Pet) string {
	var message string

	for i, pet := range pets {
		message += fmt.Sprintf("*%d\\.* %s\n", i+1, pet.Name)
	}

	return message
}
