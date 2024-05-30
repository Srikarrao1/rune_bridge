package keeper

import (
	"context"

	"bridge/x/dex/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetSellOrder set a specific sellOrder in the store from its index
func (k Keeper) SetSellOrder(ctx context.Context, sellOrder types.SellOrder) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SellOrderKeyPrefix))
	b := k.cdc.MustMarshal(&sellOrder)
	store.Set(types.SellOrderKey(
		sellOrder.Index,
	), b)
}

// GetSellOrder returns a sellOrder from its index
func (k Keeper) GetSellOrder(
	ctx context.Context,
	index string,

) (val types.SellOrder, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SellOrderKeyPrefix))

	b := store.Get(types.SellOrderKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSellOrder removes a sellOrder from the store
func (k Keeper) RemoveSellOrder(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SellOrderKeyPrefix))
	store.Delete(types.SellOrderKey(
		index,
	))
}

// GetAllSellOrder returns all sellOrder
func (k Keeper) GetAllSellOrder(ctx context.Context) (list []types.SellOrder) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SellOrderKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SellOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

