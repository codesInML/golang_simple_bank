package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{FromAccountID: account1.ID, ToAccountID: account2.ID, Amount: 50}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}
