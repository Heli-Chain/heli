package helichain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "helichain/testutil/keeper"
	"helichain/testutil/nullify"
	"helichain/x/helichain"
	"helichain/x/helichain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelichainKeeper(t)
	helichain.InitGenesis(ctx, *k, genesisState)
	got := helichain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
