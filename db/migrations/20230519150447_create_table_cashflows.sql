-- +goose Up
-- +goose StatementBegin
CREATE TABLE cashflows (
    id UUID PRIMARY KEY,
    book_id UUID NOT NULL,
    description TEXT NOT NULL,
    type VARCHAR NOT NULL,
    amount BIGINT NOT NULL,
    balance BIGINT NOT NULL,
    related_cashflow_id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX cashflows_book_id_idx ON cashflows (book_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cashflows;
-- +goose StatementEnd
