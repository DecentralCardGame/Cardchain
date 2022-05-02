package cli

import (
	"encoding/json"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdReportMatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-match [player-a] [player-b] [cards-a] [cards-b] [outcome]",
		Short: "Broadcast message ReportMatch",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPlayerA := args[0]
			argPlayerB := args[1]
			var argCardsA []uint64
			var argCardsB []uint64

			err = json.Unmarshal([]byte(args[2]), &argCardsA)
			if err != nil {
				return err
			}
			err = json.Unmarshal([]byte(args[3]), &argCardsB)
			if err != nil {
				return err
			}

			argOutcome := types.Outcome(types.Outcome_value[args[4]])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgReportMatch(
				clientCtx.GetFromAddress().String(),
				argPlayerA,
				argPlayerB,
				argCardsA,
				argCardsB,
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
