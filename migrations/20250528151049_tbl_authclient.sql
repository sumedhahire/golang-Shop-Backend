-- +goose Up
-- +goose StatementBegin

CREATE TABLE if not exists authclient (
    `id_ulid` varchar(40) ,
    `Client_uuid` varchar(255) NOT NULL,
    `Client_secret` varchar(255) NOT NULL,
    `Grant_type` varchar(255) NOT NULL,
    `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `redirect_uri` varchar(255) DEFAULT NULL,
    `UpdatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `DeletedAt` timestamp NULL DEFAULT NULL,
    `Ip` varbinary(16) DEFAULT NULL,
    `public` tinyint(1) DEFAULT NULL,
    `domain` varchar(40) not null default 'domain',
    `UserAgent` varchar(255) DEFAULT NULL,
    `CreatedBy` varchar(255) DEFAULT NULL,
    `UpdatedBy` varchar(255) DEFAULT NULL,
    `DeletedBy` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`Client_uuid`)
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
