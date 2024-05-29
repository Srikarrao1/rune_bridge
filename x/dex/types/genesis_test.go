package types_test

import (
	"testing"

	"bridge/x/dex/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated sellOrder",
			genState: &types.GenesisState{
				SellOrderList: []types.SellOrder{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated buyOrder",
			genState: &types.GenesisState{
				BuyOrderList: []types.BuyOrder{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated denomTrace",
			genState: &types.GenesisState{
				DenomTraceList: []types.DenomTrace{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
