-- +goose Up
-- +goose StatementBegin
ALTER TABLE loms_order RENAME TO loms_order;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE loms_order RENAME TO loms_order;
-- +goose StatementEnd
