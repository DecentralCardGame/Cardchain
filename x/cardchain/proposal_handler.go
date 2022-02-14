package cardchain

import (
	"sort"

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

	activeCollectionsIds := k.GetActiveCollections(ctx)
	var activeCollections []sortStruct
	if len(activeCollections) >= 3 {
		for _, id := range activeCollectionsIds {
			var collection = k.GetCollection(ctx, id)
			activeCollections = append(activeCollections, sortStruct{id, collection})
		}
		sort.SliceStable(activeCollections, func(i, j int) bool {
    		return activeCollections[i].Collection.TimeStamp < activeCollections[j].Collection.TimeStamp
			},
		)
		yeetStruct := activeCollections[0]
		yeetStruct.Collection.Status = types.CStatus_archived
		k.SetCollection(ctx, yeetStruct.Id, yeetStruct.Collection)
	}

	collection.Status = types.CStatus_active
	collection.TimeStamp = ctx.BlockHeight()

	k.SetCollection(ctx, p.CollectionId, collection)

	return nil
}

type sortStruct struct {
	Id uint64
	Collection types.Collection
}
