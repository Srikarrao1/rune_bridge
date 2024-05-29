package keeper

import (
	"bridge/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
