package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/edarha/simplebank/util"
	"github.com/stretchr/testify/assert"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, account)

	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.Equal(t, arg.Currency, account.Currency)

	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc := createRandomAccount(t)

	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, acc2)

	assert.Equal(t, acc.ID, acc2.ID)
	assert.Equal(t, acc.Owner, acc2.Owner)
	assert.Equal(t, acc.Balance, acc2.Balance)
	assert.Equal(t, acc.Currency, acc2.Currency)
	assert.WithinDuration(t, acc.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc := createRandomAccount(t)

	arg := UpdateOneAccountParams{
		ID:      acc.ID,
		Balance: util.RandomBalance(),
	}

	acc2, err := testQueries.UpdateOneAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, acc2)

	assert.Equal(t, acc.ID, acc2.ID)
	assert.Equal(t, acc.Owner, acc2.Owner)
	assert.Equal(t, arg.Balance, acc2.Balance)
	assert.Equal(t, acc.Currency, acc2.Currency)
	assert.WithinDuration(t, acc.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	assert.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)

	assert.Error(t, err)
	assert.Empty(t, acc2)
	assert.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)

	assert.NoError(t, err)
	for _, account := range accounts {
		assert.NotEmpty(t, account)
	}
}
