-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    bookname VARCHAR(255) NOT NULL,
    writer VARCHAR(255) NOT NULL,
    user_id INT REFERENCES users(id),
    deadline DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
