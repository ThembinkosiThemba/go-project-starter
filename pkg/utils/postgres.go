package utils

import (
	"context"
	"database/sql"
	"fmt"
)

func BeginTxP(ctx context.Context, db *sql.DB, query string) (*sql.Stmt, *sql.Tx, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("failed to prepare statement: %v", err)
	}

	return stmt, tx, nil
}

func PrepareContext(ctx context.Context, db *sql.DB, query string) (*sql.Rows, error) {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %v", err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all records: %v", err)
	}

	return rows, nil
}
