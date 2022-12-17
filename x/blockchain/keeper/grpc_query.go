package keeper

import (
	"github/chancebet/blockchain/x/blockchain/types"
)

var _ types.QueryServer = Keeper{}
