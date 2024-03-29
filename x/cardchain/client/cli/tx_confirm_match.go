package cli

import (
	"encoding/json"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdConfirmMatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-match [match-id] [outcome] [votes]",
		Short: "Broadcast message confirmMatch",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMatchId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argOutcome := types.Outcome(types.Outcome_value[args[1]])
			var argVotes []*types.SingleVote
			err = json.Unmarshal([]byte(args[2]), &argVotes)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgConfirmMatch(
				clientCtx.GetFromAddress().String(),
				argMatchId,
				argVotes,
				argOutcome,
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
