package keeper

import (
	"context"

	"bridge/x/dex/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BuyOrderAll(ctx context.Context, req *types.QueryAllBuyOrderRequest) (*types.QueryAllBuyOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var buyOrders []types.BuyOrder

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	buyOrderStore := prefix.NewStore(store, types.KeyPrefix(types.BuyOrderKeyPrefix))

	pageRes, err := query.Paginate(buyOrderStore, req.Pagination, func(key []byte, value []byte) error {
		var buyOrder types.BuyOrder
		if err := k.cdc.Unmarshal(value, &buyOrder); err != nil {
			return err
		}

		buyOrders = append(buyOrders, buyOrder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBuyOrderResponse{BuyOrder: buyOrders, Pagination: pageRes}, nil
}

func (k Keeper) BuyOrder(ctx context.Context, req *types.QueryGetBuyOrderRequest) (*types.QueryGetBuyOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetBuyOrder(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBuyOrderResponse{BuyOrder: val}, nil
}
