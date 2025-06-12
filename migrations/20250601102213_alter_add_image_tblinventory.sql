-- +goose Up
-- +goose StatementBegin
ALTER TABLE Tbl_Inventory
    ADD COLUMN ImageLink varchar(100) default null ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
