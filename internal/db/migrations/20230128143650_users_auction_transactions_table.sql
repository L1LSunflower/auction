-- +goose Up
-- +goose StatementBegin
create table auction_transactions (
    user_id varchar(36) not null,
    amount float not null,
    type varchar(20) not null,
    created_at timestamp not null,
    constraint foreign key (user_id) references users(id),
    index (user_id),
    index (user_id, type)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auction_transactions cascade;
-- +goose StatementEnd