package cli

import (
	"strconv"
	"strings"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitEarlyAccessProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "early-access-proposal [deposit] [description] [link] [users...]",
		Short: "Propose EarlyAccessProposal",
		Args:  cobra.MinimumNArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDeposit := args[0]
			argDescription := args[1]
			argLink := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			users := args[3:]

			for _, addr := range users {
				_, err = sdk.AccAddressFromBech32(addr)
				if err != nil {
					return err
				}
			}

			from := clientCtx.GetFromAddress()
			proposal := types.EarlyAccessProposal{
				Title:       "Early access proposal for: " + strings.Join(users, ", "),
				Description: argDescription,
				Link:        argLink,
				Users:       users,
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
