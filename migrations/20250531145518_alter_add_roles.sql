-- +goose Up
-- +goose StatementBegin
ALTER TABLE tbl_users
    ADD COLUMN role ENUM('user', 'admin') NOT NULL DEFAULT 'user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
