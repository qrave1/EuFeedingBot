-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pets
(
    id         VARCHAR(36) PRIMARY KEY,
    chat_id    INTEGER      NOT NULL,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_pets_chat_id ON pets (chat_id);

CREATE TABLE IF NOT EXISTS feeding
(
    id         VARCHAR(36) PRIMARY KEY,
    pet_id     VARCHAR(36) NOT NULL,
    feeding_at DATE        NOT NULL,
    food_type  VARCHAR(255),
    FOREIGN KEY (pet_id) REFERENCES pets (id) ON DELETE CASCADE
);
CREATE INDEX idx_feeding_pet_id ON feeding (pet_id);
CREATE INDEX idx_feedings_time ON feeding (feeding_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pets;
DROP TABLE feeding;
-- +goose StatementEnd
