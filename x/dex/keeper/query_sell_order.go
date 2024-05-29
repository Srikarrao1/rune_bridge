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

func (k Keeper) SellOrderAll(ctx context.Context, req *types.QueryAllSellOrderRequest) (*types.QueryAllSellOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sellOrders []types.SellOrder

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	sellOrderStore := prefix.NewStore(store, types.KeyPrefix(types.SellOrderKeyPrefix))

	pageRes, err := query.Paginate(sellOrderStore, req.Pagination, func(key []byte, value []byte) error {
		var sellOrder types.SellOrder
		if err := k.cdc.Unmarshal(value, &sellOrder); err != nil {
			return err
		}

		sellOrders = append(sellOrders, sellOrder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSellOrderResponse{SellOrder: sellOrders, Pagination: pageRes}, nil
}

func (k Keeper) SellOrder(ctx context.Context, req *types.QueryGetSellOrderRequest) (*types.QueryGetSellOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetSellOrder(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSellOrderResponse{SellOrder: val}, nil
}
