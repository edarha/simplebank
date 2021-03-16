package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferTX(t *testing.T) {
	store := NewStore(testDB)

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	// run n concurrency transfer transaction
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTXResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTX(context.Background(), TransferTXParam{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check result
	for i := 0; i < n; i++ {
		err := <-errs
		assert.NoError(t, err)

		result := <-results
		assert.NotEmpty(t, result)

		// account
		transfer := result.Transfer
		assert.NotEmpty(t, transfer)
		assert.Equal(t, acc1.ID, transfer.FromAccountID)
		assert.Equal(t, acc2.ID, transfer.ToAccountID)
		assert.Equal(t, amount, transfer.Amount)

		assert.NotZero(t, transfer.ID)
		assert.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		assert.NoError(t, err)
		// from entry
		fromEntry := result.FromEntry
		assert.NotEmpty(t, fromEntry)
		assert.Equal(t, acc1.ID, fromEntry.AccountID)
		assert.Equal(t, -amount, fromEntry.Amount)
		assert.NotZero(t, fromEntry.ID)
		assert.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		assert.NoError(t, err)

		// to entry
		toEntry := result.ToEntry
		assert.NotEmpty(t, toEntry)
		assert.Equal(t, acc2.ID, toEntry.AccountID)
		assert.Equal(t, amount, toEntry.Amount)
		assert.NotZero(t, toEntry.ID)
		assert.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		assert.NoError(t, err)

		// TODO: check account
	}
}
