package repository

import (
	"EuFeeding/internal/domain/entity"
	v1 "EuFeeding/internal/repository/v1"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type AnimalRepository interface {
	Add(a entity.Animal) error
	List(chatID int64) ([]entity.Animal, error)
}

type AnimalRepo struct {
	db *sqlx.DB
}

func NewAnimalRepo(db *sqlx.DB) *AnimalRepo {
	return &AnimalRepo{db: db}
}

// TODO: заменить на миграции
func initDB(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS animals (
		id TEXT PRIMARY KEY,
		chat_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_animals_chat_id ON animals(chat_id);
	`
	_, err := db.Exec(schema)
	return err
}

func (ar *AnimalRepo) Add(a entity.Animal) error {
	dto := v1.AnimalDTO{
		ID:        a.ID,
		ChatID:    a.ChatID,
		Name:      a.Name,
		CreatedAt: a.CreatedAt,
	}

	query := `
	INSERT INTO animals (id, chat_id, name, created_at)
	VALUES (:id, :chat_id, :name, :created_at)
	`
	_, err := ar.db.NamedExec(query, dto)
	return err
}

func (ar *AnimalRepo) List(chatID int64) ([]entity.Animal, error) {
	var animals []v1.AnimalDTO
	query := `SELECT * FROM animals WHERE chat_id = ? ORDER BY created_at DESC`
	err := ar.db.Select(&animals, query, chatID)

	resAnimals := make([]entity.Animal, 0, len(animals))

	for _, animal := range animals {
		resAnimals = append(resAnimals,
			entity.Animal{
				ID:        animal.ID,
				ChatID:    animal.ChatID,
				Name:      animal.Name,
				CreatedAt: animal.CreatedAt,
			},
		)
	}

	return resAnimals, err

}
