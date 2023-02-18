-- +goose Up
-- +goose StatementBegin
create table auction_comments (
    user_id varchar(36) not null,
    auction_id bigint not null,
    comment text null,
    constraint foreign key (user_id) references users(id),
    constraint foreign key (auction_id) references auction.auction(id),
    index (auction_id),
    index (user_id, auction_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auction_comments cascade;
-- +goose StatementEnd