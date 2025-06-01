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
		Use:   "encounter-create [name] [drawlist] [parameters] [image]",
		Short: "Broadcast message EncounterCreate",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var reqDrawlist []uint64
			err = json.Unmarshal([]byte(args[1]), &reqDrawlist)
			if err != nil {
				return err
			}

			var reqParameters []*types.Parameter
			err = json.Unmarshal([]byte(args[2]), &reqParameters)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEncounterCreate(
				clientCtx.GetFromAddress().String(),
				args[0],
				reqDrawlist,
				reqParameters,
				[]byte(args[3]),
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
