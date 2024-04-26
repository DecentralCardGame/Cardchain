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
		Use:   "q-cards [owner] [statuses] [card-types] [classes] [rarities] [sort-by] [name-contains] [keywords-contains] [notes-contains] [startercards-only] [balance-achors-only]",
		Short: "Query qCards",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwner := args[0]
			var reqOnlyStarterCards bool
			var reqOnlyBalanceAnchors bool
			var statuses []string
			var reqStatuses []types.Status

			err = getJsonArg(args[1], &statuses)
			if err != nil {
				return err
			}

			for _, status := range statuses {
				s, ok := types.Status_value[status]
				if !ok {
					return fmt.Errorf("invalid status %s", status)
				}
				reqStatuses = append(reqStatuses, types.Status(s))
			}

			var cardTypes []string
			var reqCardTypes []types.CardType

			err = getJsonArg(args[2], &cardTypes)
			if err != nil {
				return err
			}

			for _, cardType := range cardTypes {
				s, ok := types.CardType_value[cardType]
				if !ok {
					return fmt.Errorf("invalid class %s", cardType)
				}
				reqCardTypes = append(reqCardTypes, types.CardType(s))
			}

			var classes []string
			var reqClasses []types.CardClass

			err = getJsonArg(args[3], &classes)
			if err != nil {
				return err
			}

			for _, class := range classes {
				s, ok := types.CardClass_value[class]
				if !ok {
					return fmt.Errorf("invalid class %s", class)
				}
				reqClasses = append(reqClasses, types.CardClass(s))
			}

			var rarities []string
			var reqRarities []types.CardRarity

			err = getJsonArg(args[4], &rarities)
			if err != nil {
				return err
			}

			for _, rarity := range rarities {
				s, ok := types.CardRarity_value[rarity]
				if !ok {
					return fmt.Errorf("invalid rarity %s", rarity)
				}
				reqRarities = append(reqRarities, types.CardRarity(s))
			}

			reqSortBy := args[5]
			reqNameContains := args[6]
			reqKeywordsContains := args[7]
			reqNotesContains := args[8]

			switch args[9] {
			case "yes":
				reqOnlyStarterCards = true
			case "no":
				reqOnlyStarterCards = false
			default:
				return fmt.Errorf("arg 'only-startercard' has to be either yes or no but not %s", args[8])
			}

			switch args[10] {
			case "yes":
				reqOnlyBalanceAnchors = true
			case "no":
				reqOnlyBalanceAnchors = false
			default:
				return fmt.Errorf("arg 'balance-achors-only' has to be either yes or no but not %s", args[8])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQCardsRequest{

				Owner:              reqOwner,
				Statuses:           reqStatuses,
				CardTypes:          reqCardTypes,
				Classes:            reqClasses,
				Rarities:           reqRarities,
				SortBy:             reqSortBy,
				NameContains:       reqNameContains,
				KeywordsContains:   reqKeywordsContains,
				NotesContains:      reqNotesContains,
				OnlyStarterCard:    reqOnlyStarterCards,
				OnlyBalanceAnchors: reqOnlyBalanceAnchors,
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
