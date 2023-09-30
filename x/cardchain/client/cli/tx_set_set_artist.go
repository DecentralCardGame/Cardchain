package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetSetArtist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-set-artist [set-id] [artist]",
		Short: "Broadcast message setSetArtist",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSetId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argArtist := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetSetArtist(
				clientCtx.GetFromAddress().String(),
				argSetId,
				argArtist,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
