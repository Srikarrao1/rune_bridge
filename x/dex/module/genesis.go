package dex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"bridge/x/dex/keeper"
	"bridge/x/dex/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the sellOrder
	for _, elem := range genState.SellOrderList {
		k.SetSellOrder(ctx, elem)
	}
	// Set all the buyOrder
	for _, elem := range genState.BuyOrderList {
		k.SetBuyOrder(ctx, elem)
	}
	// Set all the denomTrace
	for _, elem := range genState.DenomTraceList {
		k.SetDenomTrace(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.SellOrderList = k.GetAllSellOrder(ctx)
	genesis.BuyOrderList = k.GetAllBuyOrder(ctx)
	genesis.DenomTraceList = k.GetAllDenomTrace(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
