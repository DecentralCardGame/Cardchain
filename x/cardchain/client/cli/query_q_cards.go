package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQCards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-cards [owner] [status] [card-type] [classes] [sort-by] [name-contains] [keywords-contains] [notes-contains]",
		Short: "Query qCards",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwner := args[0]
			var reqStatus types.Status
			if args[1] == "" {
				reqStatus = types.Status_none
			} else {
				reqStatus = types.Status(types.Status_value[args[1]])
			}
			reqCardType := args[2]
			reqClasses := args[3]
			reqSortBy := args[4]
			reqNameContains := args[5]
			reqKeywordsContains := args[6]
			reqNotesContains := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQCardsRequest{

				Owner:            reqOwner,
				Status:           reqStatus,
				CardType:         reqCardType,
				Classes:          reqClasses,
				SortBy:           reqSortBy,
				NameContains:     reqNameContains,
				KeywordsContains: reqKeywordsContains,
				NotesContains:    reqNotesContains,
			}

			res, err := queryClient.QCards(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
