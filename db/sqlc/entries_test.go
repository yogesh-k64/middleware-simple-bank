package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yogesh-k64/middleware-simple-bank/utils"
)

func createRandomEntry(t *testing.T, account Account) (Account, Entry) {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return account, entry
}

func TestCreateEntries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	_, entry1 := createRandomEntry(t, account)

	testEntry, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, testEntry)

	require.Equal(t, entry1.ID, testEntry.ID)
	require.Equal(t, entry1.AccountID, testEntry.AccountID)
	require.Equal(t, entry1.Amount, testEntry.Amount)
	require.WithinDuration(t, entry1.CreatedAt, testEntry.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntryParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}
}
