-- +goose Up
-- +goose StatementBegin
create table auction_visitors (
auction_id bigint not null,
user_id varchar(36) not null,
first_name varchar(100) null,
last_name varchar(100) null,
phone varchar(20) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auction_visitors;
-- +goose StatementEnd