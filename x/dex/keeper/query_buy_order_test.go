package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "bridge/testutil/keeper"
	"bridge/testutil/nullify"
	"bridge/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBuyOrderQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	msgs := createNBuyOrder(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetBuyOrderRequest
		response *types.QueryGetBuyOrderResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBuyOrderRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetBuyOrderResponse{BuyOrder: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBuyOrderRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetBuyOrderResponse{BuyOrder: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBuyOrderRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BuyOrder(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestBuyOrderQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	msgs := createNBuyOrder(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBuyOrderRequest {
		return &types.QueryAllBuyOrderRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BuyOrderAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BuyOrder), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BuyOrder),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BuyOrderAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BuyOrder), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BuyOrder),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.BuyOrderAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.BuyOrder),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.BuyOrderAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
