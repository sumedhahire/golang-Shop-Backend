-- +goose Up
-- +goose StatementBegin
create table if not exists Tbl_Tag(
    Id_uuid varchar(40) primary key ,
    Name varchar(20) not null ,
    Description text not null ,

    Is_Active bigint not null default 1,

    Created_at timestamp default NOW(),
    Updated_at timestamp not null ,
    Deleted_at timestamp null

);
-- +goose StatementEnd

-- +goose StatementBegin
create table if not exists Tbl_Inventory(
    Id_uuid varchar(40) primary key ,
    Name varchar(20) unique not null ,
    Description Text not null ,
    Price float not null ,

    Is_Active bigint not null default 1,

    Created_at timestamp default NOW(),
    Updated_at timestamp not null ,
    Deleted_at timestamp null

);

-- +goose StatementEnd

-- +goose StatementBegin
create table if not exists Tbl_Inventory_Tag(
    Id_uuid varchar(40) primary key ,
    InventoryId varchar(40) not null ,
    TagId varchar(40) not null ,

    Created_at timestamp default NOW(),
    Updated_at timestamp not null ,
    Deleted_at timestamp null,

    FOREIGN KEY (InventoryId) REFERENCES Tbl_Inventory(id_uuid),
    FOREIGN KEY (TagId) REFERENCES Tbl_Tag(Id_uuid)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
