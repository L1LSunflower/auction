-- +goose Up
-- +goose StatementBegin
create table users_messages (
    id bigint primary key auto_increment not null,
    sender_id varchar(36) not null,
    recipient_id varchar(36) not null,
    message text null,
    constraint foreign key (sender_id) references users (id),
    constraint foreign key (recipient_id) references users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users_messages cascade;
-- +goose StatementEnd