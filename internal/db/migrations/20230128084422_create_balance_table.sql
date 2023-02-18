-- +goose Up
-- +goose StatementBegin
create table balance (
    id varchar(36) not null,
    balance float not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp null,
    index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table balance cascade;
-- +goose StatementEnd