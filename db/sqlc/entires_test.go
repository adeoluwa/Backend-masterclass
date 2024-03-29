package backend_masterclass

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/adeoluwa/simplebank/util"
)

func createRandomEntry(t *testing.T, account Account) Entry{
	arg := CreateEnteryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entries, err := testQueries.CreateEntery(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Equal(t, arg.AccountID, entries.AccountID)
	require.Equal(t, arg.Amount, entries.Amount)

	require.NotZero(t, entries.ID)
	require.NotZero(t, entries.CreatedAt)

	return entries
}

func TestCreateEntry(t *testing.T){
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T){
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T){
	account := createRandomAccount(t)
	
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams {
		AccountID:account.ID,
		Limit:5,
		Offset:5,
	}

	entry, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entry, 5)

	for _, entry := range entry{
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
