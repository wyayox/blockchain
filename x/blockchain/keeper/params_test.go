package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github/chancebet/blockchain/testutil/keeper"
	"github/chancebet/blockchain/x/blockchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
