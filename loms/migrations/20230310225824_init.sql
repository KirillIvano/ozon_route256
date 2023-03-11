-- +goose Up
-- +goose StatementBegin
create table if not exists warehouse (
    warehouse_id bigserial primary key
);

create table if not exists warehouse_items (
    warehouse_id int8 not null,
    sku int8 not null,
    count int4 not null,

    primary key (warehouse_id, sku)
);

create table if not exists reservation (
    warehouse_id int8 not null,
    order_id int8 not null,
    sku int8 not null,
    count int4 not null,

    primary key (warehouse_id, order_id, sku)
);

create table if not exists user_order (
    order_id bigserial primary key,
    user_id int8 not null,
    order_status text not null,
    created_at timestamp not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin    

drop table if exists user_order;
drop table if exists reservation;
drop table if exists warehouse_items;
drop table if exists warehouse;

-- +goose StatementEnd
