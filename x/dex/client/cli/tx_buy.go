package cli

import (
	"strconv"

	"bridge/x/dex/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/ibc-go/v8/modules/core/04-channel/client/utils"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

// CmdSendBuy() returns the Buy send packet command.
// This command does not use AutoCLI because it gives a better UX to do not.
func CmdSendBuy() *cobra.Command {
	flagPacketTimeoutTimestamp := "packet-timeout-timestamp"

	cmd := &cobra.Command{
		Use:   "send-buy [src-port] [src-channel] [amount-denom] [amount] [price-denom] [price]",
		Short: "Send a buy over IBC",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argAmountDenom := args[2]
			argAmount, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argPriceDenom := args[4]
			argPrice, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendBuy(creator, srcPort, srcChannel, timeoutTimestamp, argAmountDenom, argAmount, argPriceDenom, argPrice)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
