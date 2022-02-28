package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCommitCouncilResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-council-response [councilId] [response] [secret]",
		Short: "Broadcast message CommitCouncilResponse",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCouncilId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			hashStringResponse := keeper.GetResponseHash(types.Response(types.Response_value[args[1]]), args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCommitCouncilResponse(
				clientCtx.GetFromAddress().String(),
				hashStringResponse,
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
