-- +goose Up
-- +goose StatementBegin
ALTER TABLE loms_order ADD COLUMN updated_at timestamp not null default now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE loms_order DROP COLUMN updated_at;
-- +goose StatementEnd
