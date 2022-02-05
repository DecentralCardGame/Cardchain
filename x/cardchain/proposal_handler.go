package cardchain

import (
	// "fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
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

	card := k.GetCard(ctx, p.CardId)

	card.Image = []byte{}
	card.Artist = card.Owner
	card.Status = types.Status_suspended

	k.SetCard(ctx, p.CardId, card)

	return nil
}
