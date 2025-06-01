package util

import (
	"context"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"inventory/ent/entgen"
	"inventory/internal/common/response"
)

func BindAndValidate(rq interface{}, e echo.Context) error {
	err := e.Bind(rq)
	if err != nil {
		return err
	}

	v := validator.New()
	err = v.Struct(rq)
	if err != nil {
		return err
	}

	return nil
}

func ExecTx(ctx context.Context, client *entgen.Client, fn func(tx *entgen.Tx) error) error {
	// Start a new transaction.
	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting tx: %w", err)
	}

	// Ensure a rollback in case of a panic.
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	// Execute the transactional function.
	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	// Commit the transaction if no error occurred.
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing tx: %w", err)
	}

	return nil
}

func GetUuid() string {
	return uuid.NewString()
}

func ConvertToResponse(data interface{}) response.BaseRS {
	return response.BaseRS{
		APIVersion: "",
		Data:       data,
		Error:      nil,
	}
}

func ConvertToListResponse(data interface{}) response.BaseRS {
	var base response.BaseRS
	base.Data = response.ItemsRs{
		Items: data,
	}

	return base
}
