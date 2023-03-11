-- +goose Up
-- +goose StatementBegin
alter table auctions add column (price float);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table auctions drop column price;
-- +goose StatementEnd