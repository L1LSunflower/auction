-- +goose Up
-- +goose StatementBegin
create table auctions (
    id bigint primary key auto_increment not null,
    category varchar(100) not null,
    owner_id varchar(36) not null,
    winner_id varchar(36) null,
    item_id bigint not null,
    short_description text null,
    start_price float null,
    minimal_price float not null,
    status varchar(10) not null,
    started_at timestamp null,
    ended_at timestamp null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp null,
    constraint foreign key (owner_id) references users(id),
    constraint foreign key (winner_id) references users(id),
    constraint foreign key (item_id) references items(id),
    index (owner_id),
    index (item_id),
    index (minimal_price),
    index (created_at),
    index (created_at, minimal_price),
    index (started_at),
    index (started_at, ended_at)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auctions cascade;
-- +goose StatementEnd