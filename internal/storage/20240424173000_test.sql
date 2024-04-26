-- +goose Up
-- +goose StatementBegin;
CREATE TABLE source (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
    url VARCHAR(255) NOT NULL
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
