package keeper

import (
	"bridge/x/dex/types"
)

var _ types.QueryServer = Keeper{}
