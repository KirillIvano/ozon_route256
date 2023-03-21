-- +goose Up
-- +goose StatementBegin
create table if not exists order_items (
    order_id int8 not null,
    sku int8 not null,
    count int4 not null,

    primary key (order_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists order_items;
-- +goose StatementEnd
