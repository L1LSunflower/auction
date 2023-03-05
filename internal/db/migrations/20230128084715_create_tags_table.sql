-- +goose Up
-- +goose StatementBegin
create table tags (
    id bigint primary key auto_increment not null,
    name varchar(255) not null,
    index (id),
    index (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tags cascade;
-- +goose StatementEnd