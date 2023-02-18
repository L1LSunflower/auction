-- +goose Up
-- +goose StatementBegin
create table items (
    id bigint primary key auto_increment not null,
    user_id varchar(36) not null,
    category_id int not null,
    name varchar(255) not null,
    tags json null,
    images json null,
    description text,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp null,
    foreign key (user_id) references users(id),
    foreign key (category_id) references categories(id),
    index (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table items cascade;
-- +goose StatementEnd