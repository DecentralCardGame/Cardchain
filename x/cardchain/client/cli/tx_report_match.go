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
		Use:   "report-match [match-id] [cards-a] [cards-b] [outcome]",
		Short: "Broadcast message ReportMatch",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMatchId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			var argCardsA []uint64
			var argCardsB []uint64

			err = json.Unmarshal([]byte(args[1]), &argCardsA)
			if err != nil {
				return err
			}
			err = json.Unmarshal([]byte(args[2]), &argCardsB)
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
				uint64(argMatchId),
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
