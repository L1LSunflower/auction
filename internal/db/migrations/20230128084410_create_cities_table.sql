-- +goose Up
-- +goose StatementBegin
create table cities (
    id int primary key not null,
    name varchar(100) not null,
    is_active tinyint not null,
    created_at timestamp not null,
    updated_at timestamp null,
    deleted_at timestamp null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table cities cascade;
-- +goose StatementEnd