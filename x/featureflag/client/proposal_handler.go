package client

import (
	"github.com/DecentralCardGame/Cardchain/x/featureflag/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.CmdSubmitFlagEnableProposal)
