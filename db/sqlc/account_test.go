package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yogesh-k64/middleware-simple-bank/utils"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	// Create a new account
	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	testAccount, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, testAccount)

	require.Equal(t, account1.ID, testAccount.ID)
	require.Equal(t, account1.Owner, testAccount.Owner)
	require.Equal(t, account1.Balance, testAccount.Balance)
	require.Equal(t, account1.Currency, testAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, testAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account1.ID, updatedAccount.ID)
	require.Equal(t, account1.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account1.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	testAccount, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, testAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
