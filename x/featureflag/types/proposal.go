package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// ProposalTypeChange defines the type for a ParameterChangeProposal
	ProposalTypeFlagEnable = "FeatureFlagEnable"
)

func (c *FlagEnableProposal) ProposalRoute() string { return RouterKey }

func (c *FlagEnableProposal) ProposalType() string { return ProposalTypeFlagEnable }

func (c *FlagEnableProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(c)
	if err != nil {
		return err
	}
	// TODO More validation
	return nil
}

func init() {
	govtypes.RegisterProposalType(ProposalTypeFlagEnable)
}
