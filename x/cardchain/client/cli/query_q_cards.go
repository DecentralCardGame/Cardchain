package cli

import (
	"fmt"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQCards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-cards [owner] [status] [card-type] [classes] [sort-by] [name-contains] [keywords-contains] [notes-contains] [only-startercard]",
		Short: "Query qCards",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwner := args[0]
			var reqOnlyStarterCard bool
			var reqStatus types.QueryQCardsRequest_Status
			if args[1] == "" {
				reqStatus = types.QueryQCardsRequest_none
			} else {
				reqStatus = types.QueryQCardsRequest_Status(types.QueryQCardsRequest_Status_value[args[1]])
			}
			reqCardType := args[2]
			reqClasses := args[3]
			reqSortBy := args[4]
			reqNameContains := args[5]
			reqKeywordsContains := args[6]
			reqNotesContains := args[7]

			switch args[8]{
			case "yes":
				reqOnlyStarterCard = true
			case "no":
				reqOnlyStarterCard = false
			default:
				return fmt.Errorf("arg 'only-startercard' has to be either yes or no but not %s", args[8])
			}

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
				OnlyStarterCard:  reqOnlyStarterCard,
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
