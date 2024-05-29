package dex_test

import (
	"testing"

	keepertest "bridge/testutil/keeper"
	"bridge/testutil/nullify"
	dex "bridge/x/dex/module"
	"bridge/x/dex/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		SellOrderList: []types.SellOrder{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		BuyOrderList: []types.BuyOrder{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		DenomTraceList: []types.DenomTrace{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, k, genesisState)
	got := dex.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.SellOrderList, got.SellOrderList)
	require.ElementsMatch(t, genesisState.BuyOrderList, got.BuyOrderList)
	require.ElementsMatch(t, genesisState.DenomTraceList, got.DenomTraceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
