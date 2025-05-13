package repository

import (
	"github.com/qrave1/PetFeedingBot/internal/domain/entity"
	"github.com/qrave1/PetFeedingBot/internal/repository/dto/v1"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type PetRepository interface {
	Add(pet entity.Pet) error
	List(chatID int64) ([]entity.Pet, error)
}

type PetRepo struct {
	db *sqlx.DB
}

func NewPetRepo(db *sqlx.DB) *PetRepo {
	return &PetRepo{db: db}
}

func (ar *PetRepo) Add(pet entity.Pet) error {
	dto := v1.Pet{
		ID:        pet.ID,
		ChatID:    pet.ChatID,
		Name:      pet.Name,
		CreatedAt: pet.CreatedAt,
	}

	query := `
	INSERT INTO pets (id, chat_id, name, created_at)
	VALUES (:id, :chat_id, :name, :created_at)
	`
	_, err := ar.db.NamedExec(query, dto)

	return err
}

func (ar *PetRepo) List(chatID int64) ([]entity.Pet, error) {
	var pets []v1.Pet

	query := `SELECT * FROM pets WHERE chat_id = ? ORDER BY created_at`

	err := ar.db.Select(&pets, query, chatID)

	resPets := make([]entity.Pet, 0, len(pets))

	for _, pet := range pets {
		resPets = append(resPets,
			entity.Pet{
				ID:        pet.ID,
				ChatID:    pet.ChatID,
				Name:      pet.Name,
				CreatedAt: pet.CreatedAt,
			},
		)
	}

	return resPets, err

}
