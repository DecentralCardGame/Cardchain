package cli

import (
	"encoding/json"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/keywords"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSaveCardContent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-card-content [card-id] [content] [notes] [artist]",
		Short: "Broadcast message SaveCardContent",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			cardobj, err := keywords.Unmarshal([]byte(args[1]))
			if err != nil {
				return err
			}

			cardbytes, err := json.Marshal(cardobj)
			if err != nil {
				return err
			}

			argNotes := args[2]
			argArtist := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSaveCardContent(
				clientCtx.GetFromAddress().String(),
				argCardId,
				cardbytes,
				// argImage,
				// argFullArt,
				argNotes,
				argArtist,
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
