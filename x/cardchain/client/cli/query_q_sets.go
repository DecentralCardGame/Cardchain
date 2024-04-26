package cli

import (
	"encoding/json"
	"slices"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQSets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-sets [status] [contributors] [contains-cards] [owner]",
		Short: "Query q_sets",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var (
				reqStatus    types.CStatus
				ignoreStatus bool = false
			)

			if !slices.Contains([]string{"\"\"", "''"}, args[0]) {
				reqStatus = types.CStatus(types.CStatus_value[args[0]])
			} else {
				ignoreStatus = true
			}

			var argContributors []string
			err = json.Unmarshal([]byte(args[1]), &argContributors)
			if err != nil {
				return err
			}

			var argContainsCards []uint64
			err = json.Unmarshal([]byte(args[2]), &argContainsCards)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQSetsRequest{

				Status:        reqStatus,
				IgnoreStatus:  ignoreStatus,
				Contributors:  argContributors,
				ContainsCards: argContainsCards,
				Owner:         args[3],
			}

			res, err := queryClient.QSets(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
