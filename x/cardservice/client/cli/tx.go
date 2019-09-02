package cli

import (
	//"fmt"
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/types"

	"github.com/DecentralCardGame/cardobject"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	cardserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Cardservice transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cardserviceTxCmd.AddCommand(client.PostCommands(
		GetCmdBuyCardScheme(cdc),
		GetCmdSaveCardContent(cdc),
		GetCmdVoteCard(cdc),
		GetCmdTransferCard(cdc),
		GetCmdDonateToCard(cdc),
		GetCmdCreateUser(cdc),
	)...)

	return cardserviceTxCmd
}

// GetCmdBuyCardScheme is the CLI command for sending a BuyCardScheme transaction
func GetCmdBuyCardScheme(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy-card-scheme [amount]",
		Short: "bid for a card scheme",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			/*if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}*/

			coins, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyCardScheme(coins, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSaveCardContent is the CLI command for saving the content of a card
func GetCmdSaveCardContent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "save-card-content [card id] [content]",
		Short: "save the content of a card",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			/*if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}*/

			cardobj, err := cardobject.ProcessCard(args[1])
			if err != nil {
				return err
			}
			if strings.Compare(args[1], cardobj) != 0 {
				return errors.New("Card content string and card content object don't match. This means ProcessCard() removed or could not parse parts of the card content.")
			}

			cardId, err := strconv.ParseUint(args[0], 10, 64);
			if err != nil {
				return err
			}

			msg := types.NewMsgSaveCardContent(cardId, []byte(cardobj), cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdVoteCard is the CLI command for voting for a card
func GetCmdVoteCard(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vote-card [card id] [vote type]",
		Short: "vote a card as inappropriate, underpowered, overpowered or fair enough",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			/*if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}*/

			cardId, err := strconv.ParseUint(args[0], 10, 64);
			if err != nil {
				return err
			}

			msg := types.NewMsgVoteCard(cardId, args[1], cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdTransferCard is the CLI command for transferring the ownership of a card
func GetCmdTransferCard(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transfer-card [card id] [receiver address]",
		Short: "transfer the ownership of a card",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			/*if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}*/

			cardId, err := strconv.ParseUint(args[0], 10, 64);
			if err != nil {
				return err
			}

			receiver, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferCard(cardId, cliCtx.GetFromAddress(), receiver)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDonateToCard is the CLI command for donating credits to a card
func GetCmdDonateToCard(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "donate-to-card [card id] [amount]",
		Short: "donate credits to the voters pool of a card",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			/*if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}*/

			cardId, err := strconv.ParseUint(args[0], 10, 64);
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDonateToCard(cardId, cliCtx.GetFromAddress(), coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdCreateUser is the CLI command for voting for a card
func GetCmdCreateUser(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-user [Addresss] [alias]",
		Short: "create a user, this means giving starting credits and starting cards",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			/*if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}*/

			newUser, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUser(cliCtx.GetFromAddress(), newUser, args[1])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
