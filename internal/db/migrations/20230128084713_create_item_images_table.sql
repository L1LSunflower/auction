-- +goose Up
-- +goose StatementBegin
create table item_images (
    item_id bigint not null,
    filename varchar(255) not null,
    constraint foreign key (item_id) references items(id),
    index (item_id),
    index (filename),
    index (item_id, filename)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table item_images cascade;
-- +goose StatementEnd