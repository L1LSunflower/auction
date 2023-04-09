-- +goose Up
-- +goose StatementBegin
alter table auctions add column (visit_status varchar(20) null, visit_start_date datetime null, visit_end_date datetime null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table auctions drop column visit_status, drop column visit_start_date, drop column visit_end_date;
-- +goose StatementEnd