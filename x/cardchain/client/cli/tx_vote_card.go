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

func CmdVoteCard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-card [card-id] [vote-type]",
		Short: "Broadcast message VoteCard",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argVoteType, found := types.VoteType_value[args[1]]
			if !found {
				return fmt.Errorf("rarity has to be in %s", maps.Keys(types.CardRarity_value))
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVoteCard(
				clientCtx.GetFromAddress().String(),
				argCardId,
				types.VoteType(argVoteType),
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
