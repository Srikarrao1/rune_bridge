package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:         PortID,
		SellOrderList:  []SellOrder{},
		BuyOrderList:   []BuyOrder{},
		DenomTraceList: []DenomTrace{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in sellOrder
	sellOrderIndexMap := make(map[string]struct{})

	for _, elem := range gs.SellOrderList {
		index := string(SellOrderKey(elem.Index))
		if _, ok := sellOrderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sellOrder")
		}
		sellOrderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in buyOrder
	buyOrderIndexMap := make(map[string]struct{})

	for _, elem := range gs.BuyOrderList {
		index := string(BuyOrderKey(elem.Index))
		if _, ok := buyOrderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for buyOrder")
		}
		buyOrderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in denomTrace
	denomTraceIndexMap := make(map[string]struct{})

	for _, elem := range gs.DenomTraceList {
		index := string(DenomTraceKey(elem.Index))
		if _, ok := denomTraceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for denomTrace")
		}
		denomTraceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
