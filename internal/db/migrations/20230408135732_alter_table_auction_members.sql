-- +goose Up
-- +goose StatementBegin
alter table auction_members add column (price float not null, first_name varchar(100) null, last_name varchar(100) null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table auction_members drop column price, drop column first_name, drop column last_name;
-- +goose StatementEnd