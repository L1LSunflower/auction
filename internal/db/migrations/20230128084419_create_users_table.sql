-- +goose Up
-- +goose StatementBegin
create table users (
    id varchar(36) primary key not null,
    phone varchar(20) not null,
    email varchar(150) not null,
    password varchar(255) not null,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    city varchar(50) null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp not null,
    constraint email_phone_unique unique (email, phone),
    index (id),
    index (phone),
    index (email),
    index (id, created_at),
    index (email, created_at)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users cascade;
-- +goose StatementEnd