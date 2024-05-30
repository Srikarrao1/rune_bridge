package keeper

import (
	"bridge/x/bridge/types"
	"fmt"
	"strings"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	// bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

func isIBCToken(denom string) bool {
	return strings.HasPrefix(denom, "ibc/")
}

func(k Keeper) SafeBurn(ctx sdk.Context, port string, channel string, sender sdk.AccAddress, denom string, amount int32) error {
	if isIBCToken(denom) {
		//Burn the tokens
		if err := k.BurnTokens(ctx, sender, sdk.NewCoin(denom, sdkmath.NewInt(int64(amount)))); err != nil {
			return err
		}
	} else {
		if err := k.LockTokens(ctx, port, channel, sender, sdk.NewCoin(denom, sdkmath.NewInt(int64(amount)))); err != nil {
			return err
		}
	}
	return nil
}


func(k Keeper) BurnTokens(ctx sdk.Context, sender sdk.AccAddress, tokens sdk.Coin) error {
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(tokens)); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		panic(fmt.Sprintf("cannot burn coin after a successful send to module account"))
	}
	return nil
}

func(k Keeper) LockTokens(ctx sdk.Context, sourcePort string, sourceChannel string, sender sdk.AccAddress, tokens sdk.Coin) error {
	//Create the escrow address for the address

	escrowAddress := ibctransfertypes.GetEscrowAddress(sourcePort, sourceChannel)
	    // escrow source tokens. It fails if balance insufficient

		if err  := k.bankKeeper.SendCoins(
			ctx, sender, escrowAddress, sdk.NewCoins(tokens),
		); err != nil {
			return err
		}
		return nil

	
}