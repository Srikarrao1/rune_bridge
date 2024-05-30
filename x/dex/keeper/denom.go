package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	sdkmath "cosmossdk.io/math"
	"bridge/x/dex/types"
)

func(k Keeper) SaveVoucherDenom(ctx sdk.Context, port string, channel string, denom string) {
	voucher := k.VoucherDenom(port, channel, denom)

	//Store the original denom

	_, saved := k.GetDenomTrace(ctx, voucher)
	if !saved {
		k.SetDenomTrace(ctx, types.DenomTrace{
			Index: voucher,
			Port: port,
			Channel: channel,
			Origin: denom,
		})
	}
}

func (k Keeper) VoucherDenom(port string, channel string, denom string) string {
	sourcePrefix := ibctransfertypes.GetDenomPrefix(port, channel)

	prefixedDenom := sourcePrefix + denom

	denomTrace := ibctransfertypes.ParseDenomTrace(prefixedDenom)
	voucher := denomTrace.IBCDenom()
	return voucher[:16]
}

func (k Keeper) OriginalDenom(ctx sdk.Context, port string, channel string, voucher string) (string, bool) {
	trace, exist := k.GetDenomTrace(ctx, voucher)

	if !exist {
		if trace.Port == port && trace.Channel == channel {
			return trace.Origin, true
		}
	}
	return "", true
}

func (k Keeper) SafeMint(ctx sdk.Context, port string, channel string, receiver sdk.AccAddress, denom string, amount int32) error {
	if isIBCToken(denom) {
		//Mint the IBC tokens
		if err := k.MintTokens(ctx, receiver, sdk.NewCoin(denom, sdkmath.NewInt(int64(amount)))); err != nil {
			return err
		}
	} else {
		//Unlock the native runes
		if err := k.UnlockTokens(
			ctx,
			port,
			channel,
			receiver,
			sdk.NewCoin(denom, sdkmath.NewInt(int64(amount))),
		); err != nil {
			return err
		}
	}
	return nil
} 

func (k Keeper) MintTokens(ctx sdk.Context, receiver sdk.AccAddress, tokens sdk.Coin ) error {
	// mint the tokens

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(tokens)); err != nil {
		return err
	}
	//send to receiver 
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiver, sdk.NewCoins(tokens), 
	); err != nil {
		panic(fmt.Sprintf("unable to send coins from module to account despite previously minting coins to module account: %v", err))
	}
	return nil
}

func (k Keeper) UnlockTokens(ctx sdk.Context, sourcePort string, sourceChannel string, receiver sdk.AccAddress, tokens sdk.Coin) error {
	//unlock the tokens

	//create an escrow address to hold the tokens
	escrowAddress := ibctransfertypes.GetEscrowAddress(sourcePort, sourceChannel)
	//escrow source tokens 

	if err := k.bankKeeper.SendCoins(ctx, escrowAddress, receiver, sdk.NewCoins(tokens)); err != nil {
		return err
	}
	return nil
}

