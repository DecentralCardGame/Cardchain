package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateuser())
	cmd.AddCommand(CmdBuyCardScheme())
	cmd.AddCommand(CmdVoteCard())
	cmd.AddCommand(CmdSaveCardContent())
	cmd.AddCommand(CmdTransferCard())
	cmd.AddCommand(CmdDonateToCard())
	cmd.AddCommand(CmdAddArtwork())
	cmd.AddCommand(CmdSubmitCopyrightProposal())
	cmd.AddCommand(CmdChangeArtist())
	cmd.AddCommand(CmdRegisterForCouncil())
	cmd.AddCommand(CmdReportMatch())
	cmd.AddCommand(CmdSubmitMatchReporterProposal())
	cmd.AddCommand(CmdApointMatchReporter())
	cmd.AddCommand(CmdCreateCollection())
	cmd.AddCommand(CmdAddCardToCollection())
	cmd.AddCommand(CmdFinalizeCollection())
	cmd.AddCommand(CmdBuyCollection())
	cmd.AddCommand(CmdRemoveCardFromCollection())
	cmd.AddCommand(CmdRemoveContributorFromCollection())
	cmd.AddCommand(CmdAddContributorToCollection())
	cmd.AddCommand(CmdSubmitCollectionProposal())
	cmd.AddCommand(CmdCreateSellOffer())
	cmd.AddCommand(CmdBuyCard())
	cmd.AddCommand(CmdRemoveSellOffer())
	cmd.AddCommand(CmdAddArtworkToCollection())
	// this line is used by starport scaffolding # 1

	return cmd
}
