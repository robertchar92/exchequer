-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    username VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX books_user_id_idx ON books (user_id);
CREATE INDEX books_username_idx ON books (username);
CREATE INDEX books_name_idx ON books (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
