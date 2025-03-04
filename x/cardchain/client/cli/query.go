package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group cardchain queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQCard())

	cmd.AddCommand(CmdQCardContent())

	cmd.AddCommand(CmdQUser())

	cmd.AddCommand(CmdQCardchainInfo())

	cmd.AddCommand(CmdQVotingResults())

	cmd.AddCommand(CmdQCards())

	cmd.AddCommand(CmdQMatch())

	cmd.AddCommand(CmdQSet())

	cmd.AddCommand(CmdQSellOffer())

	cmd.AddCommand(CmdQCouncil())

	cmd.AddCommand(CmdQMatches())

	cmd.AddCommand(CmdQSellOffers())

	cmd.AddCommand(CmdQServer())

	cmd.AddCommand(CmdQSets())

	cmd.AddCommand(CmdRarityDistribution())

	cmd.AddCommand(CmdQCardContents())

	cmd.AddCommand(CmdQAccountFromZealy())

	cmd.AddCommand(CmdQEncounters())

	cmd.AddCommand(CmdQEncounter())

	cmd.AddCommand(CmdQEncountersWithImage())

	cmd.AddCommand(CmdQEncounterWithImage())

	// this line is used by starport scaffolding # 1

	return cmd
}
