package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dzwiedz90/helpdesk-service-users/model"
	"github.com/dzwiedz90/helpdesk-service-users/pkg/database"
	"github.com/dzwiedz90/helpdesk-service-users/service/serviceconfig"
)

var (
	defaultTransactionOptions = &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	}
)

func NewCore() Core {
	return Core{}
}

type Core struct{}

func (c *Core) CreateUser(ctx context.Context, cfg *serviceconfig.ServerConfig, creq *model.CreateUser) (int64, error) {
	tx, err := cfg.DB.SQL.BeginTx(ctx, defaultTransactionOptions)
	if err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to start transaction: %v", err))
		return 0, err
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to rollback transaction: %v", rollbackErr))
			// error log with context "Failed to rollback transaction"
		}
	}()

	res, err := database.InsertUser(ctx, tx, creq, cfg)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to commit transaction: %v", err))
		return 0, err
	}

	return res, nil
}

func (c *Core) GetUser(ctx context.Context, cfg *serviceconfig.ServerConfig, id int64) (*model.User, error) {
	tx, err := cfg.DB.SQL.BeginTx(ctx, defaultTransactionOptions)
	if err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to start transaction: %v", err))
		return nil, err
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to rollback transaction: %v", rollbackErr))
			// error log with context "Failed to rollback transaction"
		}
	}()

	res, err := database.GetUser(ctx, tx, id, cfg)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to commit transaction: %v", err))
		return nil, err
	}

	return res, nil
}

func (c *Core) GetAlUsers(ctx context.Context, cfg *serviceconfig.ServerConfig) ([]*model.User, error) {
	tx, err := cfg.DB.SQL.BeginTx(ctx, defaultTransactionOptions)
	if err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to start transaction: %v", err))
		return nil, err
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to rollback transaction: %v", rollbackErr))
			// error log with context "Failed to rollback transaction"
		}
	}()

	res, err := database.GetAllUsers(ctx, tx, cfg)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to commit transaction: %v", err))
		return nil, err
	}

	return res, nil
}

func (c *Core) UpdateUser(ctx context.Context, cfg *serviceconfig.ServerConfig, creq *model.UpdateUser) error {
	tx, err := cfg.DB.SQL.BeginTx(ctx, defaultTransactionOptions)
	if err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to start transaction: %v", err))
		return err
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to rollback transaction: %v", rollbackErr))
			// error log with context "Failed to rollback transaction"
		}
	}()

	err = database.UpdateUser(ctx, tx, creq, cfg)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to commit transaction: %v", err))
		return err
	}

	return nil
}

func (c *Core) DeleteUser(ctx context.Context, cfg *serviceconfig.ServerConfig, id int64) error {
	tx, err := cfg.DB.SQL.BeginTx(ctx, defaultTransactionOptions)
	if err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to start transaction: %v", err))
		return err
	}

	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to rollback transaction: %v", rollbackErr))
			// error log with context "Failed to rollback transaction"
		}
	}()

	err = database.DeleteUser(ctx, tx, id, cfg)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		cfg.Logger.ErrorLogger(fmt.Sprintf("Failed to commit transaction: %v", err))
		return err
	}

	return nil
}
