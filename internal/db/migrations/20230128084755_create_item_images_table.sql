-- +goose Up
-- +goose StatementBegin
create table item_images (
    id bigint not null,
    path varchar(255) not null,
    constraint foreign key (id) references items(id),
    index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table item_images cascade;
-- +goose StatementEnd