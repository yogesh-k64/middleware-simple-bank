package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yogesh-k64/middleware-simple-bank/utils"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {

	arg := CreatetransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.Createtransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}
func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1, account2)
}
