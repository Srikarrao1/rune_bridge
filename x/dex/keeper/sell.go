package keeper

import (
	"errors"

	"bridge/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// TransmitSellPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitSellPacket(
	ctx sdk.Context,
	packetData types.SellPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.ScopedKeeper().GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, errorsmod.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %s", err)
	}

	return k.ibcKeeperFn().ChannelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvSellPacket processes packet reception
func (k Keeper) OnRecvSellPacket(ctx sdk.Context, packet channeltypes.Packet, data types.SellPacketData) (packetAck types.SellPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	return packetAck, nil
}

// OnAcknowledgementSellPacket responds to the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementSellPacket(ctx sdk.Context, packet channeltypes.Packet, data types.SellPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// In case of error we mint back the native token
        receiver, err := sdk.AccAddressFromBech32(data.Seller)
        if err != nil {
            return err
        }

        if err := k.SafeMint(ctx, packet.SourcePort, packet.SourceChannel, receiver, data.AmountDenom, data.Amount); err != nil {
            return err
        }

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.SellPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}
		// Get the sell order book
        pairIndex := types.OrderBookIndex(packet.SourcePort, packet.SourceChannel, data.AmountDenom, data.PriceDenom)
        book, found := k.GetSellOrder(ctx, pairIndex)
        if !found {
            panic("sell order book must exist")
        }

		 // Append the remaining amount of the order
		 if packetAck.RemainingAmount > 0 {
            _, err := book.AppendOrder(data.Seller, packetAck.RemainingAmount, data.Price)
            if err != nil {
                return err
            }

            // Save the new order book
            k.SetSellOrder(ctx, book)
        }

        // Mint the gains
        if packetAck.Gain > 0 {
            receiver, err := sdk.AccAddressFromBech32(data.Seller)
            if err != nil {
                return err
            }

            finalPriceDenom, saved := k.OriginalDenom(ctx, packet.SourcePort, packet.SourceChannel, data.PriceDenom)
            if !saved {
                // If it was not from this chain we use voucher as denom
                finalPriceDenom = k.VoucherDenom(packet.DestinationPort, packet.DestinationChannel, data.PriceDenom)
            }

            if err := k.SafeMint(ctx, packet.SourcePort, packet.SourceChannel, receiver, finalPriceDenom, packetAck.Gain); err != nil {
                return err
            }
        }


		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutSellPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutSellPacket(ctx sdk.Context, packet channeltypes.Packet, data types.SellPacketData) error {

	  // In case of error we mint back the native token
	  receiver, err := sdk.AccAddressFromBech32(data.Seller)
	  if err != nil {
		  return err
	  }
  
	  if err := k.SafeMint(ctx, packet.SourcePort, packet.SourceChannel, receiver, data.AmountDenom, data.Amount); err != nil {
		  return err
	  }
	  
	return nil
}
