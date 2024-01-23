package cli

import (
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitFlagEnableProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flag-enable-proposal [module] [name] [deposit]",
		Short: "Propose FlagEnableProposal",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDeposit := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			proposal := types.FlagEnableProposal{
				Title:       "Request to enable featureflag `" + args[1] + "` for module `" + args[0] + "`",
				Description: "See `" + args[0] + "-" + args[1] + "` for more information on the set.",
				Module:      args[0],
				Name:        args[1],
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
