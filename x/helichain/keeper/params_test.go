package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "helichain/testutil/keeper"
	"helichain/x/helichain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HelichainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
