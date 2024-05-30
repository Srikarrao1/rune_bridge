package keeper

import (
	"context"
	"errors"
	"bridge/x/dex/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

func (k msgServer) SendBuy(goCtx context.Context, msg *types.MsgSendBuy) (*types.MsgSendBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
    // Cannot send a order if the pair doesn't exist
    pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.AmountDenom, msg.PriceDenom)
    _, found := k.GetBuyOrder(ctx, pairIndex)
    if !found {
        return &types.MsgSendBuyResponse{}, errors.New("the pair doesn't exist")
    }

    // Lock the token to send
    sender, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        return &types.MsgSendBuyResponse{}, err
    }

    // Use SafeBurn to ensure no new native tokens are minted
    if err := k.SafeBurn(ctx, msg.Port, msg.ChannelID, sender, msg.PriceDenom, msg.Amount*msg.Price); err != nil {
        return &types.MsgSendBuyResponse{}, err
    }

    // Save the voucher received on the other chain, to have the ability to resolve it into the original denom
    k.SaveVoucherDenom(ctx, msg.Port, msg.ChannelID, msg.PriceDenom)

	var packet types.BuyPacketData

	packet.Buyer = msg.Creator
	packet.AmountDenom = msg.AmountDenom
	packet.Amount = msg.Amount
	packet.PriceDenom = msg.PriceDenom
	packet.Price = msg.Price

	// Transmit the packet
	_, err = k.TransmitBuyPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	);
	if err != nil {
		return nil, err
	}


	return &types.MsgSendBuyResponse{}, nil
}
