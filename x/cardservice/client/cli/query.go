package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	cardserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the cardservice module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cardserviceQueryCmd.AddCommand(flags.GetCommands(
		GetCmdCard(storeKey, cdc),
		GetCmdUser(storeKey, cdc),
		GetCmdCardList(storeKey, cdc),
		GetCmdVotableCardList(storeKey, cdc),
		GetCmdCardchainInfo(storeKey, cdc),
	)...)

	return cardserviceQueryCmd
}

// GetCmdCard queries information about a card
func GetCmdCard(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "card [cardId]",
		Short: "queries info of a card with given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/card/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not query the card with id - %s \n", string(id))
				return nil
			}

			//fmt.Println(string(res))
			//return nil

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdUser queries information about a user
func GetCmdUser(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "user [address]",
		Short: "Query all info of user, that is beyond the pure bank account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/user/%s", queryRoute, address), nil)
			if err != nil {
				fmt.Printf("could not query user - %s \n", string(address))
				return nil
			}

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdCardList queries a list of 50 or all? cards
func GetCmdCardList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "cards",
		Short: "cards",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cards", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query cards\n")
				return nil
			}

			//fmt.Println(string(res))
			//return nil

			//var out types.QueryResResolve
			//cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(string(res))
		},
	}
}

// GetCmdVotableCardList looks up the cards votable by a user
func GetCmdVotableCardList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "votable-cards [address]",
		Short: "Query cards votable of a user.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/votable-cards/%s", queryRoute, address), nil)
			if err != nil {
				fmt.Printf("could not get query votable cards\n")
				return nil
			}

			//fmt.Println(string(res))
			//return nil

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdUser queries information about a user
func GetCmdCardchainInfo(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "cardchain-info",
		Short: "Query all info of the cardchain that is relevant for the cards",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cardchain-info", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not query cardchain info\n")
				return nil
			}

			//fmt.Println(string(res))
			//return nil

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
