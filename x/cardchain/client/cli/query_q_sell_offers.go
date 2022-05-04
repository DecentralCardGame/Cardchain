package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQSellOffers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "q-sell-offers [price-down] [price-up] [seller] [buyer] [card] [status]",
		Short: "Query qSellOffers",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ignore := types.NewIgnoreSellOffers()
			reqPriceDown := "0ucredits"
			reqPriceUp := "0ucredits"
			var (
				reqCard uint64
			)

			if args[0] != "" && args[1] != "" {
				reqPriceDownRaw, err := sdk.ParseCoinNormalized(args[0])
				if err != nil {
					return err
				}
				reqPriceUpRaw, err := sdk.ParseCoinNormalized(args[1])
				if err != nil {
					return err
				}
				reqPriceDown = reqPriceDownRaw.String()
				reqPriceUp = reqPriceUpRaw.String()
			}

			reqSeller := args[2]
			reqBuyer := args[3]

			if args[4] == "" {
				ignore.Card = true
			} else {
				reqCard, err = cast.ToUint64E(args[4])
				if err != nil {
					return err
				}
			}

			if args[5] == "" {
				ignore.Status = true
			}
			reqStatus := types.SellOfferStatus(types.SellOfferStatus_value[args[5]])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQSellOffersRequest{

				PriceDown: reqPriceDown,
				PriceUp:   reqPriceUp,
				Seller:    reqSeller,
				Buyer:     reqBuyer,
				Card:      reqCard,
				Status:    reqStatus,
				Ignore:    &ignore,
			}

			res, err := queryClient.QSellOffers(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
