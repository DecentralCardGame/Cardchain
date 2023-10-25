package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitCopyrightProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copyright-proposal [card-id] [deposit] [description] [link]",
		Short: "Propose CopyrightProposal",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argDeposit := args[1]
			argDescription := args[2]
			argLink := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			proposal := types.CopyrightProposal{
				Title:       "Copyright violation `" + args[0] + "`",
				Description: argDescription,
				Link:        argLink,
				CardId:      argCardId,
			}

			deposit, err := sdk.ParseCoinsNormalized(argDeposit)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&proposal, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
