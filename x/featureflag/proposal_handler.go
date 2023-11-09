package featureflag

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/keeper"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// NewProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.FlagEnableProposal:
			return handleFlagEnableProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(errors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}

func handleFlagEnableProposal(ctx sdk.Context, k keeper.Keeper, p *types.FlagEnableProposal) error {
	err := k.FlagExists(p.Module, p.Name)
	if err != nil {
		return err
	}
	
	key := keeper.GetKey(p.Module, p.Name)
	
	flag := k.GetFlag(ctx, key)
	flag.Set = true
	k.SetFlag(ctx, key, flag)
	
	return nil
}
