package cardchain

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CopyrightProposal:
			return handleCopyrightProposal(ctx, k, c)
		case *types.MatchReporterProposal:
			return handleMatchReporterProposal(ctx, k, c)
		case *types.CollectionProposal:
			return handleCollectionProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}

func handleMatchReporterProposal(ctx sdk.Context, k keeper.Keeper, p *types.MatchReporterProposal) error {
	return k.ApointMatchReporter(ctx, p.Reporter)
}

func handleCopyrightProposal(ctx sdk.Context, k keeper.Keeper, p *types.CopyrightProposal) error {
	card := k.GetCard(ctx, p.CardId)

	card.Image = []byte{}
	card.Artist = card.Owner
	card.Status = types.Status_suspended

	k.SetCard(ctx, p.CardId, card)

	return nil
}

func handleCollectionProposal(ctx sdk.Context, k keeper.Keeper, p *types.CollectionProposal) error {
	collection := k.GetCollection(ctx, p.CollectionId)

	if collection.Status != types.CStatus_finalized {
		return sdkerrors.Wrapf(types.ErrCollectionNotInDesign, "Collection status is %s but should be finalized", collection.Status)
	}

	collection.Status = types.CStatus_active
	collection.ExpireBlock = ctx.BlockHeight() + 1032000  // 3 months I guess

	k.SetCollection(ctx, p.CollectionId, collection)

	return nil
}
