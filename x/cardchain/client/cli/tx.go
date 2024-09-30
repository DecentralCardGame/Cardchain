package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
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
	cmd.AddCommand(CmdChangeArtist())
	cmd.AddCommand(CmdRegisterForCouncil())
	cmd.AddCommand(CmdReportMatch())
	cmd.AddCommand(CmdApointMatchReporter())
	cmd.AddCommand(CmdCreateSet())
	cmd.AddCommand(CmdAddCardToSet())
	cmd.AddCommand(CmdFinalizeSet())
	cmd.AddCommand(CmdBuyBoosterPack())
	cmd.AddCommand(CmdRemoveCardFromSet())
	cmd.AddCommand(CmdRemoveContributorFromSet())
	cmd.AddCommand(CmdAddContributorToSet())
	cmd.AddCommand(CmdCreateSellOffer())
	cmd.AddCommand(CmdBuyCard())
	cmd.AddCommand(CmdRemoveSellOffer())
	cmd.AddCommand(CmdAddArtworkToSet())
	cmd.AddCommand(CmdAddStoryToSet())
	cmd.AddCommand(CmdSetCardRarity())
	cmd.AddCommand(CmdCreateCouncil())
	cmd.AddCommand(CmdCommitCouncilResponse())
	cmd.AddCommand(CmdRevealCouncilResponse())
	cmd.AddCommand(CmdRestartCouncil())
	cmd.AddCommand(CmdRewokeCouncilRegistration())
	cmd.AddCommand(CmdConfirmMatch())
	cmd.AddCommand(CmdSetProfileCard())
	cmd.AddCommand(CmdOpenBoosterPack())
	cmd.AddCommand(CmdTransferBoosterPack())
	cmd.AddCommand(CmdSetSetStoryWriter())
	cmd.AddCommand(CmdSetSetArtist())
	cmd.AddCommand(CmdSetUserWebsite())
	cmd.AddCommand(CmdSetUserBiography())
	cmd.AddCommand(CmdMultiVoteCard())
	cmd.AddCommand(CmdMsgOpenMatch())
	cmd.AddCommand(CmdSetSetName())
	cmd.AddCommand(CmdChangeAlias())
	cmd.AddCommand(CmdInviteEarlyAccess())
	cmd.AddCommand(CmdDisinviteEarlyAccess())
	cmd.AddCommand(CmdConnectZealyAccount())
	cmd.AddCommand(CmdEncounterCreate())
	cmd.AddCommand(CmdEncounterDo())
	cmd.AddCommand(CmdEncounterClose())
	// this line is used by starport scaffolding # 1

	return cmd
}
