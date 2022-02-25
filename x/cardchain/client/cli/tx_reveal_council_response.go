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

func CmdRevealCouncilResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reveal-council-response [councilId] [response] [secret]",
		Short: "Broadcast message RevealCouncilResponse",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCouncilId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argResponse := types.Response(types.Response_value[args[1]])
			argSecret := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevealCouncilResponse(
				clientCtx.GetFromAddress().String(),
				argResponse,
				argSecret,
				argCouncilId,
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
