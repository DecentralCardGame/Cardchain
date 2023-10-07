package cli

import (
	"fmt"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

var _ = strconv.Itoa(0)

func CmdSetCardRarity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-card-rarity [card-id] [set-id] [rarity]",
		Short: "Broadcast message SetCardRarity",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argSetId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			rar, found := types.CardRarity_value[args[2]]
			if (!found) {
				return fmt.Errorf("Rarity has to be in %s", maps.Keys(types.CardRarity_value))
			}

			argRarity := types.CardRarity(rar)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetCardRarity(
				clientCtx.GetFromAddress().String(),
				argCardId,
				argSetId,
				argRarity,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
