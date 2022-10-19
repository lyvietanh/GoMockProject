package db

import (
	"database/sql"
)

type Store struct {
	Queries *Queries
	Db      *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Db:      db,
		Queries: New(db),
	}
}

// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		rollbackErr := tx.Rollback()
// 		if rollbackErr != nil {
// 			return fmt.Errorf("tx error: %v, rollback err: %v", err, rollbackErr)
// 		}
// 	}
// 	return tx.Commit()
// }

// type TransferTxParams struct {
// }

// type TransferTxResult struct {
// }

// func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
// 	var result TransferTxResult

// }
