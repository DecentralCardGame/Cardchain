package cli

import (
    "strconv"
	
	 "encoding/json"
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

var _ = strconv.Itoa(0)

func CmdCardVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card-vote [vote]",
		Short: "Broadcast message CardVote",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argVote := new(types.SingleVote)
					err = json.Unmarshal([]byte(args[0]), argVote)
    				if err != nil {
                		return err
            		}
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCardVote(
				clientCtx.GetFromAddress().String(),
				argVote,
				
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