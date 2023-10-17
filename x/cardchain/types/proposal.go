package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// Module name for minting coins
const CoinsIssuerName = "cardchain"

const (
	// ProposalTypeChange defines the type for a ParameterChangeProposal
	ProposalTypeCopyright     = "Copyright"
	ProposalTypeMatchReporter = "MatchReporter"
	ProposalTypeSet           = "Set"
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

func (c *SetProposal) ProposalRoute() string { return RouterKey }

func (c *SetProposal) ProposalType() string { return ProposalTypeSet }

func (c *SetProposal) ValidateBasic() error {
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
	govtypes.RegisterProposalType(ProposalTypeSet)
}
