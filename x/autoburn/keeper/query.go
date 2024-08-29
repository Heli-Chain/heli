package keeper

import (
	"helichain/x/autoburn/types"
)

var _ types.QueryServer = Keeper{}
