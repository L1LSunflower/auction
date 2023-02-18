-- +goose Up
-- +goose StatementBegin
create table auction.auction (
    id bigint primary key auto_increment not null,
    owner_id varchar(36) not null,
    winner_id varchar(36) null,
    item_id bigint not null,
    title varchar(255) not null,
    description text not null,
    start_price float null,
    minimal_price float not null,
    auction_period int not null,
    is_active tinyint null,
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
    index (title),
    index (minimal_price),
    index (auction_period),
    index (created_at),
    index (created_at, minimal_price, auction_period)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auction.auction cascade;
-- +goose StatementEnd