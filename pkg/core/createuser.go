package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dzwiedz90/helpdesk-service-users/driver"
	"github.com/dzwiedz90/helpdesk-service-users/logs"
	"github.com/dzwiedz90/helpdesk-service-users/model"
	"github.com/dzwiedz90/helpdesk-service-users/pkg/database"
)

var (
	defaultTransactionOptions = &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	}
)

func CreateUser(ctx context.Context, db *driver.DB, creq *model.CreateUser) (int64, error) {
	tx, err := db.SQL.BeginTx(ctx, defaultTransactionOptions)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to start transaction: %v", err))
		return 0, err
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			logs.ErrorLogger(fmt.Sprintf("Failed to rollback transaction: %v", rollbackErr))
			// error log with context "Failed to rollback transaction"
		}
	}()

	res, err := database.InsertUser(ctx, tx, creq)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to commit transaction: %v", err))
		return 0, err
	}

	return res, nil
}

func GetUser(ctx context.Context, db *driver.DB, id int64) (*model.User, error) {
	tx, err := db.SQL.BeginTx(ctx, defaultTransactionOptions)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to start transaction: %v", err))
		return nil, err
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			logs.ErrorLogger(fmt.Sprintf("Failed to rollback transaction: %v", rollbackErr))
			// error log with context "Failed to rollback transaction"
		}
	}()

	res, err := database.GetUser(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		logs.ErrorLogger(fmt.Sprintf("Failed to commit transaction: %v", err))
		return nil, err
	}

	return res, nil
}
