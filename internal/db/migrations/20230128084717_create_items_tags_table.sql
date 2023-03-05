-- +goose Up
-- +goose StatementBegin
create table item_tags (
    item_id bigint not null,
    tag_id bigint not null,
    constraint foreign key (item_id) references items(id),
    constraint foreign key (tag_id) references tags(id),
    index (item_id, tag_id),
    index (item_id),
    index (tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table item_tags cascade;
-- +goose StatementEnd