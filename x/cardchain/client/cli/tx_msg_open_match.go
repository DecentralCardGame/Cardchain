package cli

import (
	"encoding/json"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMsgOpenMatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-match [player-a] [player-b] [player-a-deck] [player-b-deck]",
		Short: "Broadcast message OpenMatch",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var argPlayerADeck []uint64
			var argPlayerBDeck []uint64
			err = json.Unmarshal([]byte(args[2]), &argPlayerADeck)
			if err != nil {
				return err
			}
			err = json.Unmarshal([]byte(args[3]), &argPlayerBDeck)
			if err != nil {
				return err
			}

			msg := types.NewMsgOpenMatch(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				argPlayerADeck,
				argPlayerBDeck,
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
