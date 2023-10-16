package cli

import (
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitMatchReporterProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "match-reporter-proposal [deposit] [description]",
		Short: "Propose MatchReporterProposal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDeposit := args[0]
			argDescription := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argReporter := clientCtx.GetFromAddress().String()

			from := clientCtx.GetFromAddress()
			proposal := types.MatchReporterProposal{
				Title:       "Match Reporter Request `" + argReporter + "`",
				Description: argDescription,
				Reporter:    argReporter,
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
