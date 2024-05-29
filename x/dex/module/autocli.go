package dex

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "bridge/api/bridge/dex"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "SellOrderAll",
					Use:       "list-sell-order",
					Short:     "List all sell-order",
				},
				{
					RpcMethod:      "SellOrder",
					Use:            "show-sell-order [id]",
					Short:          "Shows a sell-order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "BuyOrderAll",
					Use:       "list-buy-order",
					Short:     "List all buy-order",
				},
				{
					RpcMethod:      "BuyOrder",
					Use:            "show-buy-order [id]",
					Short:          "Shows a buy-order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "DenomTraceAll",
					Use:       "list-denom-trace",
					Short:     "List all denom-trace",
				},
				{
					RpcMethod:      "DenomTrace",
					Use:            "show-denom-trace [id]",
					Short:          "Shows a denom-trace",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
