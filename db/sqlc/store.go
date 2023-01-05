package db

import (
	"context"
	"database/sql"
	"fmt"
)

// 이곳에서는 Transaction의 ACID를 지키기 위한 항목들이 정리되어 있습니다.
// 예를들어 A가 B에게 송금을 하게 된다면, 최소 두개의 Tx가 발생을 하여야 하고 둘중 하나라도 실패하면 모두 실패해야 합니다.


type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db : db,
		Queries : New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err!= nil {
		if rbErr := tx.Rollback(); rbErr != nil{
			return fmt.Errorf("tx err: %v,  rb err : %v",err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// contains the input parameters of thr transfer Tx
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 	`json:"to_account_id"`
	Amount		int64	`json:"amount"`
}

// result of the Tx
type TransferTxResult struct {
	Transfer Transfer 	`json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account 	`json:"to_account"`
	FromEntry Entry 	`json:"from_entry"`
	ToEntry Entry 		`json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult , error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID : arg.FromAccountID,
			ToAccountID : arg.ToAccountID,
			Amount : arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams {
			AccountID : arg.FromAccountID,
			Amount : arg.Amount,
		})

		if err != nil{
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID : arg.ToAccountID,
			Amount : arg.Amount,
		})

		if err != nil{
			return err
		}

		// TODO : update Account Balance


		return nil
	})

	return result,err
}