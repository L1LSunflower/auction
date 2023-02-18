-- +goose Up
-- +goose StatementBegin
create table categories (
    id int primary key auto_increment not null,
    name varchar(100) not null,
    index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table categories cascade;
-- +goose StatementEnd