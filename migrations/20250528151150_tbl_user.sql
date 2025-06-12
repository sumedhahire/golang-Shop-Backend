-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists tbl_users (
    `id_uuid` varchar(40),
    `firstname` varchar(32) NOT NULL,
    `lastname` varchar(32) NOT NULL,
    `email` varchar(255) NOT NULL UNIQUE,
    `birth_date` timestamp NULL DEFAULT NULL,
    `password` varchar(255) NOT NULL,
    `is_active` int DEFAULT '1',
    `zip_code` int DEFAULT NULL,
    `address` text,
    `ip_address` varchar(40),
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id_uuid`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
