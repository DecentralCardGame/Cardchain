package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQFlag() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-flag [module] [name]",
		Short: "Query qFlag",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqModule := args[0]
			reqName := args[1]
			
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQFlagRequest{
				Module: reqModule,
				Name: reqName,
			}

			res, err := queryClient.QFlag(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
