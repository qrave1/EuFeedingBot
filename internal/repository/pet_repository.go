package repository

import (
	"EuFeeding/internal/domain/entity"
	"EuFeeding/internal/repository/dto/v1"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type PetRepository interface {
	Add(a entity.Pet) error
	List(chatID int64) ([]entity.Pet, error)
}

type PetRepo struct {
	db *sqlx.DB
}

func NewPetRepo(db *sqlx.DB) *PetRepo {
	return &PetRepo{db: db}
}

// TODO: заменить на миграции
func initDB(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS pets (
		id TEXT PRIMARY KEY,
		chat_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_pets_chat_id ON pets(chat_id);
	`
	_, err := db.Exec(schema)
	return err
}

func (ar *PetRepo) Add(a entity.Pet) error {
	dto := v1.PetDTO{
		ID:        a.ID,
		ChatID:    a.ChatID,
		Name:      a.Name,
		CreatedAt: a.CreatedAt,
	}

	query := `
	INSERT INTO pets (id, chat_id, name, created_at)
	VALUES (:id, :chat_id, :name, :created_at)
	`
	_, err := ar.db.NamedExec(query, dto)
	return err
}

func (ar *PetRepo) List(chatID int64) ([]entity.Pet, error) {
	var pets []v1.PetDTO
	query := `SELECT * FROM pets WHERE chat_id = ? ORDER BY created_at DESC`
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
