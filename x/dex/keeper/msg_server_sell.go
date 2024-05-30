package keeper

import (
	"context"
	"errors"

	"bridge/x/dex/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	// channel "github.com/cosmos/ibc-go/v8/modules/core/04-channel"
)

func (k msgServer) SendSell(goCtx context.Context, msg *types.MsgSendSell) (*types.MsgSendSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// If an order book doesn't exist, throw an error
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.AmountDenom, msg.PriceDenom)
	_, found := k.GetSellOrder(ctx, pairIndex)
	if !found {
		return &types.MsgSendSellResponse{}, errors.New("the pair doesn't exist")
	}

	//get Sender's address
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgSendSellResponse{}, err
	}

	//use safeBurn to ensure no new native tokens are minted
	if err := k.SafeBurn(ctx, msg.Port, msg.ChannelID, sender, msg.AmountDenom, msg.Amount); err != nil {
		return &types.MsgSendSellResponse{}, err
	}

	//Save the voucher received on the other chain, to have the ability to resolve it into the original denom
	k.SaveVoucherDenom(ctx, msg.Port, msg.ChannelID, msg.AmountDenom)

	// Construct the packet
	var packet types.SellPacketData

	packet.Seller = msg.Creator
	packet.AmountDenom = msg.AmountDenom
	packet.Amount = msg.Amount
	packet.PriceDenom = msg.PriceDenom
	packet.Price = msg.Price

	// Transmit the packet
	_, err = k.TransmitSellPacket(
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

	return &types.MsgSendSellResponse{}, nil
}

// ...

func (k Keeper) OnRecvSellOrderPacket(ctx sdk.Context, packet channeltypes.Packet, data types.SellPacketData) (packetAck types.SellPacketAck, err error) {
    if err := data.ValidateBasic(); err != nil {
        return packetAck, err
    }

    pairIndex := types.OrderBookIndex(packet.SourcePort, packet.SourceChannel, data.AmountDenom, data.PriceDenom)
    book, found := k.GetBuyOrder(ctx, pairIndex)
    if !found {
        return packetAck, errors.New("the pair doesn't exist")
    }

    // Fill sell order
    remaining, liquidated, gain, _ := book.FillSellOrder(types.Order{
        Amount: data.Amount,
        Price:  data.Price,
    })

    // Return remaining amount and gains
    packetAck.RemainingAmount = remaining.Amount
    packetAck.Gain = gain

    // Before distributing sales, we resolve the denom
    // First we check if the denom received comes from this chain originally
    finalAmountDenom, saved := k.OriginalDenom(ctx, packet.DestinationPort, packet.DestinationChannel, data.AmountDenom)
    if !saved {
        // If it was not from this chain we use voucher as denom
        finalAmountDenom = k.VoucherDenom(packet.SourcePort, packet.SourceChannel, data.AmountDenom)
    }

    // Dispatch liquidated buy orders
    for _, liquidation := range liquidated {
        liquidation := liquidation
        addr, err := sdk.AccAddressFromBech32(liquidation.Creator)
        if err != nil {
            return packetAck, err
        }

        if err := k.SafeMint(ctx, packet.DestinationPort, packet.DestinationChannel, addr, finalAmountDenom, liquidation.Amount); err != nil {
            return packetAck, err
        }
    }

    // Save the new order book
    k.SetBuyOrder(ctx, book)

    return packetAck, nil
}

