package cli

import (
	"strconv"
	"encoding/json"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/DecentralCardGame/cardobject/keywords"
)

var _ = strconv.Itoa(0)

func CmdSaveCardContent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-card-content [card-id] [content] [image] [full-art] [notes] [owner]",
		Short: "Broadcast message SaveCardContent",
		Args:  cobra.ExactArgs(6),
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

			argImage := []byte(args[2])
			argFullArt := args[3]
			argNotes := args[4]
			argOwner := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSaveCardContent(
				clientCtx.GetFromAddress().String(),
				argCardId,
				cardbytes,
				argImage,
				argFullArt,
				argNotes,
				argOwner,
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
