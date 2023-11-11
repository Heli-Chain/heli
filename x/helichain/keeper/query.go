package keeper

import (
	"helichain/x/helichain/types"
)

var _ types.QueryServer = Keeper{}
