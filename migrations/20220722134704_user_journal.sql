-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_journal (
    id SERIAL PRIMARY KEY,
    source_id uuid NOT NULL,
    target_id uuid NOT NULL,
    time timestamptz NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_journal;
-- +goose StatementEnd
