package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Module name for minting coins
const CoinsIssuerName = "cardchain"

const (
	// ProposalTypeChange defines the type for a ParameterChangeProposal
	ProposalTypeCopyright     = "Copyright"
	ProposalTypeMatchReporter = "MatchReporter"
	ProposalTypeCollection    = "Collection"
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

func (c *MatchReporterProposal) ProposalRoute() string { return RouterKey }

func (c *MatchReporterProposal) ProposalType() string { return ProposalTypeMatchReporter }

func (c *MatchReporterProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(c)
	if err != nil {
		return err
	}
	// TODO More validation
	return nil
}

func (c *CollectionProposal) ProposalRoute() string { return RouterKey }

func (c *CollectionProposal) ProposalType() string { return ProposalTypeCollection }

func (c *CollectionProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(c)
	if err != nil {
		return err
	}
	// TODO More validation
	return nil
}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCopyright)
	govtypes.RegisterProposalType(ProposalTypeMatchReporter)
	govtypes.RegisterProposalType(ProposalTypeCollection)
}
