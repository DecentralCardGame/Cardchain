package cli

import (
	"fmt"
	"strconv"
	"encoding/json"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQMatches() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-matches [timestamp-down] [timestamp-up] [contains-users] [card-played] [reporter] [outcome]",
		Short: "Query qMatches",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ignore := types.NewIgnore()
			var (
				reqTimestampDown, reqTimestampUp uint64
				reqOutcome types.Outcome
				argCards []uint64
				reqContainsUsers []string
			)

			if args[0] == "" && args[1] == "" {
				ignore.Timestamp = true
			} else {
				reqTimestampDown, err = cast.ToUint64E(args[0])
				if err != nil {
					return err
				}
				reqTimestampUp, err = cast.ToUint64E(args[1])
				if err != nil {
					return err
				}
			}

			err = json.Unmarshal([]byte(args[2]), &reqContainsUsers)
			if err != nil {
				return err
			}
			if len(reqContainsUsers) > 2 {
				return fmt.Errorf("There can only be 2 players involved in a match, not %d", len(reqContainsUsers))
			}

			err = json.Unmarshal([]byte(args[3]), &argCards)
			if err != nil {
				return err
			}

			if args[4] == "" {
				ignore.Reporter = true
			}
			reqReporter := args[4]

			if args[5] == "" {
				ignore.Outcome = true
			} else {
				reqOutcome = types.Outcome(types.Outcome_value[args[5]])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQMatchesRequest{
				TimestampDown: reqTimestampDown,
				TimestampUp:   reqTimestampUp,
				ContainsUsers: reqContainsUsers,
				CardsPlayed:   argCards,
				Reporter:			 reqReporter,
				Outcome:       reqOutcome,
				Ignore:				 &ignore,
			}

			res, err := queryClient.QMatches(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
