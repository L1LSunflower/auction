-- +goose Up
-- +goose StatementBegin
create table auction_members (
    auction_id bigint not null,
    participant_id varchar(36) not null,
    constraint foreign key (auction_id) references auction.auction(id),
    constraint foreign key (participant_id) references  users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auction_members cascade;
-- +goose StatementEnd