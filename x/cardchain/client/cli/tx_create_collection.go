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

func CmdCreateCollection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-collection [name] [artist] [story-writer] [contributors]",
		Short: "Broadcast message createCollection",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			var argContributors []string
			argArtist := args[1]
			argStoryWriter := args[2]

			err = json.Unmarshal([]byte(args[3]), &argContributors)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCollection(
				clientCtx.GetFromAddress().String(),
				argArtist,
				argStoryWriter,
				argName,
				argContributors,
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
