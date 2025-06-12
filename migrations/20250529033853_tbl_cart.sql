-- +goose Up
-- +goose StatementBegin
create table if not exists Tbl_Cart(
    Id_uuid varchar(40) primary key ,
    ProductId varchar(40) ,
    UserId varchar(40),
    Status ENUM('cart', 'canceled', 'brought') not null ,

    Created_at timestamp default NOW(),
    Updated_at timestamp not null ,
    Deleted_at timestamp null,

    FOREIGN KEY (ProductId) REFERENCES Tbl_Inventory(id_uuid),
    FOREIGN KEY (UserId) REFERENCES  tbl_users(id_uuid)


);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
