-- +goose Up
-- +goose StatementBegin
create table items (
    id bigint primary key auto_increment not null,
    user_id varchar(36) not null,
    name varchar(255) not null,
    tag1 varchar(50) null,
    tag2 varchar(50) null,
    tag3 varchar(50) null,
    tag4 varchar(50) null,
    tag5 varchar(50) null,
    tag6 varchar(50) null,
    tag7 varchar(50) null,
    tag8 varchar(50) null,
    tag9 varchar(50) null,
    tag10 varchar(50) null,
    images text null,
    description text,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp null,
    foreign key (user_id) references users(id),
    index (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table items cascade;
-- +goose StatementEnd