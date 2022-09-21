package db

import (
	"context"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreaterRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, account.ID, entry.AccountID)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}
func TestCreateEntry(t *testing.T) {
	CreaterRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := CreaterRandomEntry(t)

	enttry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, enttry2)

	require.Equal(t, entry1.ID, enttry2.ID)
	require.Equal(t, entry1.AccountID, enttry2.AccountID)
	require.Equal(t, entry1.Amount, enttry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt, enttry2.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createArg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		}
		entry, err := testQueries.CreateEntry(context.Background(), createArg)

		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry2 := range entries {
		require.NotEmpty(t, entry2)
	}
}
