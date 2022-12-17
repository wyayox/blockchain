package blockchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github/chancebet/blockchain/testutil/keeper"
	"github/chancebet/blockchain/testutil/nullify"
	"github/chancebet/blockchain/x/blockchain"
	"github/chancebet/blockchain/x/blockchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlockchainKeeper(t)
	blockchain.InitGenesis(ctx, *k, genesisState)
	got := blockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
