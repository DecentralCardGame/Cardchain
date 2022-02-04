package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Module name for minting coins
const CoinsIssuerName = "cardchain"

const (
	// ProposalTypeChange defines the type for a ParameterChangeProposal
	ProposalTypeCopyright = "Copyright"
)

func (c *CopyrightProposal) ProposalRoute() string { return RouterKey }

func (c *CopyrightProposal) ProposalType() string { return ProposalTypeCopyright }

func (c *CopyrightProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(c)
	if err != nil {
		return err
	}
	// TODO More validation
	return nil
}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCopyright)
}
