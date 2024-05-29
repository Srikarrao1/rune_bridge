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

func createNBuyOrder(keeper keeper.Keeper, ctx context.Context, n int) []types.BuyOrder {
	items := make([]types.BuyOrder, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetBuyOrder(ctx, items[i])
	}
	return items
}

func TestBuyOrderGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBuyOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBuyOrder(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBuyOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBuyOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBuyOrder(ctx,
			item.Index,
		)
		_, found := keeper.GetBuyOrder(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestBuyOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNBuyOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBuyOrder(ctx)),
	)
}
