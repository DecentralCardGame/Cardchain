package client

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the param change proposal handler.
var CopyrightProposalHandler = govclient.NewProposalHandler(cli.CmdSubmitCopyrightProposal)
var MatchReporterProposalHandler = govclient.NewProposalHandler(cli.CmdSubmitMatchReporterProposal)
var SetProposalHandler = govclient.NewProposalHandler(cli.CmdSubmitSetProposal)
var EarlyAccessProposalHandler = govclient.NewProposalHandler(cli.CmdSubmitEarlyAccessProposal)
