package keeper

import (
	"context"

	"bridge/x/dex/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetBuyOrder set a specific buyOrder in the store from its index
func (k Keeper) SetBuyOrder(ctx context.Context, buyOrder types.BuyOrder) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BuyOrderKeyPrefix))
	b := k.cdc.MustMarshal(&buyOrder)
	store.Set(types.BuyOrderKey(
		buyOrder.Index,
	), b)
}

// GetBuyOrder returns a buyOrder from its index
func (k Keeper) GetBuyOrder(
	ctx context.Context,
	index string,

) (val types.BuyOrder, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BuyOrderKeyPrefix))

	b := store.Get(types.BuyOrderKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBuyOrder removes a buyOrder from the store
func (k Keeper) RemoveBuyOrder(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BuyOrderKeyPrefix))
	store.Delete(types.BuyOrderKey(
		index,
	))
}

// GetAllBuyOrder returns all buyOrder
func (k Keeper) GetAllBuyOrder(ctx context.Context) (list []types.BuyOrder) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.BuyOrderKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BuyOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
