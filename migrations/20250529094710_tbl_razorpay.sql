-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tbl_Payment (
    Id_uuid VARCHAR(40) PRIMARY KEY,

    UserId VARCHAR(40) NOT NULL,
    InventoryId VARCHAR(40) NOT NULL,

    RazorpayOrderId VARCHAR(100) NOT NULL,
    RazorpayPaymentId VARCHAR(100),
    RazorpaySignature VARCHAR(256),

    Amount FLOAT NOT NULL,
    Status VARCHAR(50) NOT NULL DEFAULT 'created',

    Created_at TIMESTAMP DEFAULT NOW(),
    Updated_at TIMESTAMP NOT NULL,
    Deleted_at TIMESTAMP NULL,

    FOREIGN KEY (UserId) REFERENCES tbl_users(Id_uuid),
    FOREIGN KEY (InventoryId) REFERENCES Tbl_Inventory(Id_uuid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
