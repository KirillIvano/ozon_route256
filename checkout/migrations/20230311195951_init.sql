-- +goose Up
-- +goose StatementBegin
create table if not exists cart_item (
    user_id int8 not null,
    sku int4 not null,
    count int8 not null,

    primary key (user_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists cart_item;
-- +goose StatementEnd
