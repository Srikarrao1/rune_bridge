package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bridge/testutil/keeper"
	"bridge/x/bridge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BridgeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
