package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQSets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-sets [status]",
		Short: "Query q_sets",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var (
				reqStatus    types.CStatus
				ignoreStatus bool = false
			)

			if args[0] != "" {
				reqStatus = types.CStatus(types.CStatus_value[args[0]])
			} else {
				ignoreStatus = true
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQSetsRequest{

				Status:       reqStatus,
				IgnoreStatus: ignoreStatus,
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