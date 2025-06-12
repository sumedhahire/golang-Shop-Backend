-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists `authtoken` (
                                           `id_uuid` varchar(40) NOT NULL ,
                                           `auth_uuid` varchar(255) NOT NULL,
                                           `auth_xref` varchar(255) DEFAULT NULL,
                                           `accesstoken` varchar(255) DEFAULT NULL,
                                           `accesstokencreatedat` timestamp NULL DEFAULT NULL,
                                           `accesstokenexpiresin` bigint DEFAULT NULL,
                                           `clientid` varchar(255) DEFAULT NULL,
                                           `user_ulid` varchar(255) DEFAULT NULL,
                                           `refreshtoken` varchar(255) DEFAULT NULL,
                                           `refreshtokencreatedat` timestamp NULL DEFAULT NULL,
                                           `refreshtokenexpiresin` bigint DEFAULT NULL,
                                           `createdat` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                           `updatedat` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                           `deletedat` timestamp NULL DEFAULT NULL,
                                           `ip_address` varchar(255) DEFAULT NULL,
                                           `useragent` varchar(255) DEFAULT NULL,
                                           `createdby` varchar(255) DEFAULT NULL,
                                           `updatedby` varchar(255) DEFAULT NULL,
                                           `deletedby` varchar(255) DEFAULT NULL,
                                           `code` varchar(255) DEFAULT NULL,
                                           `codecreatedat` timestamp NULL DEFAULT NULL,
                                           `codeexpiresin` bigint DEFAULT NULL,
                                           `redirect_uri` varchar(255) DEFAULT NULL,
                                           `scope` varchar(255) DEFAULT NULL,
                                           `codechallenge` varchar(255) DEFAULT NULL,
                                           PRIMARY KEY (`id_uuid`),
                                           FOREIGN KEY (user_ulid) REFERENCES tbl_users(id_uuid),
                                           FOREIGN KEY (clientid) REFERENCES  authclient(Client_uuid)

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
