package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	random "github.com/nithinkm8055/simplebank/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	ctx := context.TODO()
	account := createRandomAccount(t)

	entry, err := testQueries.CreateEntry(ctx, CreateEntryParams{
		AccountID: account.ID,
		Amount:    int64(random.RandomAmount(int(account.Balance))),
	})

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	assert.Equal(t, account.ID, entry.AccountID)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	ctx := context.TODO()
	entry := createRandomEntry(t)
	require.NotEmpty(t, entry)

	t.Run("it=retrieves created entry should pass", func(t *testing.T) {
		getEntry, err := testQueries.GetEntry(ctx, entry.ID)
		require.NoError(t, err)
		require.NotEmpty(t, getEntry)

		assert.Equal(t, entry.ID, getEntry.ID)
		assert.Equal(t, entry.Amount, getEntry.Amount)
	})

	t.Run("it=retrieves non-existent entry should fail", func(t *testing.T) {
		getEntry, err := testQueries.GetEntry(ctx, uuid.New())
		require.Error(t, err)
		require.Empty(t, getEntry)
	})
}

func TestListEntries(t *testing.T) {
	numEntries := 4
	createEntries := make([]Entry, 0)
	for i := 0; i < numEntries; i++ {
		createEntries = append(createEntries, createRandomEntry(t))
	}

	getEntries, err := testQueries.ListEntries(context.TODO(), ListEntriesParams{
		Offset: 0,
		Limit:  100,
	})
	require.NoError(t, err)
	require.NotEmpty(t, getEntries)

	require.True(t, len(getEntries) >= len(createEntries))
	// validate each element of creatEntries is in getEntries
	for i := range createEntries {
		assert.Contains(t, getEntries, createEntries[i])
	}
}
