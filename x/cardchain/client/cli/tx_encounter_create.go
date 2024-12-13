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

func CmdEncounterCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encounter-create [drawlist] [parameters] [image]",
		Short: "Broadcast message EncounterCreate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var reqDrawlist []uint64
			err = json.Unmarshal([]byte(args[0]), &reqDrawlist)
			if err != nil {
				return err
			}

			var reqParameters map[string]string
			err = json.Unmarshal([]byte(args[1]), &reqParameters)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEncounterCreate(
				clientCtx.GetFromAddress().String(),
				reqDrawlist,
				reqParameters,
				[]byte(args[1]),
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
