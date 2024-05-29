package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "bridge/testutil/keeper"
	"bridge/testutil/nullify"
	"bridge/x/dex/keeper"
	"bridge/x/dex/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSellOrder(keeper keeper.Keeper, ctx context.Context, n int) []types.SellOrder {
	items := make([]types.SellOrder, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSellOrder(ctx, items[i])
	}
	return items
}

func TestSellOrderGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSellOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSellOrder(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSellOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSellOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSellOrder(ctx,
			item.Index,
		)
		_, found := keeper.GetSellOrder(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSellOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSellOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSellOrder(ctx)),
	)
}
