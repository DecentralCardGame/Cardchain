package cli

import (
	"strconv"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCardSaveContent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card-save-content [card-id] [content] [notes] [artist] [balance-anchor]",
		Short: "Broadcast message CardSaveContent",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argContent := args[1]
			argNotes := args[2]
			argArtist := args[3]
			argBalanceAnchor, err := cast.ToBoolE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCardSaveContent(
				clientCtx.GetFromAddress().String(),
				argCardId,
				argContent,
				argNotes,
				argArtist,
				argBalanceAnchor,
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
