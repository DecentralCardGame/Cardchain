package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQCardContents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-card-contents [card-ids]",
		Short: "Query q_card_contents",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var reqCardIds []uint64
			for _, id := range args[0:] {
				convId, err := cast.ToUint64E(id)
				if err != nil {
					return err
				}
				reqCardIds = append(reqCardIds, convId)
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQCardContentsRequest{

				CardIds: reqCardIds,
			}

			res, err := queryClient.QCardContents(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
