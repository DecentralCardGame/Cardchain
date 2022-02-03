package cardchain

import (
	// "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewCopyrightProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CopyrightProposal:
			return handleCopyrightProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}

func handleCopyrightProposal(ctx sdk.Context, k keeper.Keeper, p *types.CopyrightProposal) error {
	// for _, c := range p.Changes {
	// 	ss, ok := k.GetSubspace(c.Subspace)
	// 	if !ok {
	// 		return sdkerrors.Wrap(proposal.ErrUnknownSubspace, c.Subspace)
	// 	}
	//
	// 	k.Logger(ctx).Info(
	// 		fmt.Sprintf("attempt to set new parameter value; key: %s, value: %s", c.Key, c.Value),
	// 	)
	//
	// 	if err := ss.Update(ctx, []byte(c.Key), []byte(c.Value)); err != nil {
	// 		return sdkerrors.Wrapf(proposal.ErrSettingParameter, "key: %s, value: %s, err: %s", c.Key, c.Value, err.Error())
	// 	}
	// }

	return nil
}
