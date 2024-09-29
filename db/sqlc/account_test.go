package db

import (
	"context"
	"testing"

	random "github.com/nithinkm8055/simplebank/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    random.RandomOwner(6),
		Balance:  int64(random.RandomAmount(500)),
		Currency: random.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Balance, account.Balance)

	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
