package cardchain

import (
	"sort"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// NewProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CopyrightProposal:
			return handleCopyrightProposal(ctx, k, c)
		case *types.MatchReporterProposal:
			return handleMatchReporterProposal(ctx, k, c)
		case *types.SetProposal:
			return handleSetProposal(ctx, k, c)
		case *types.EarlyAccessProposal:
			return handleEarlyAccessProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(errors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}

func handleMatchReporterProposal(ctx sdk.Context, k keeper.Keeper, p *types.MatchReporterProposal) error {
	return k.SetMatchReporter(ctx, p.Reporter)
}

func handleCopyrightProposal(ctx sdk.Context, k keeper.Keeper, p *types.CopyrightProposal) error {
	card := k.Cards.Get(ctx, p.CardId)
	image := k.Images.Get(ctx, card.ImageId)

	image.Image = []byte{}
	card.Artist = card.Owner
	card.Status = types.Status_suspended

	k.SetLastCardModifiedNow(ctx)
	k.Cards.Set(ctx, p.CardId, card)
	k.Images.Set(ctx, card.ImageId, image)

	return nil
}

func handleSetProposal(ctx sdk.Context, k keeper.Keeper, p *types.SetProposal) error {
	set := k.Sets.Get(ctx, p.SetId)

	if set.Status != types.CStatus_finalized {
		return sdkerrors.Wrapf(types.ErrSetNotInDesign, "Set status is %s but should be finalized", set.Status)
	}

	activeSetsIds := k.GetActiveSets(ctx)
	var activeSets []sortStruct
	if len(activeSets) >= int(k.GetParams(ctx).ActiveSetsAmount) {
		for _, id := range activeSetsIds {
			var set = k.Sets.Get(ctx, id)
			activeSets = append(activeSets, sortStruct{id, set})
		}
		sort.SliceStable(activeSets, func(i, j int) bool {
			return activeSets[i].Set.TimeStamp < activeSets[j].Set.TimeStamp
		},
		)
		yeetStruct := activeSets[0]
		yeetStruct.Set.Status = types.CStatus_archived
		k.Sets.Set(ctx, yeetStruct.Id, yeetStruct.Set)
	}

	set.Status = types.CStatus_active
	set.TimeStamp = ctx.BlockHeight()

	k.Sets.Set(ctx, p.SetId, set)

	return nil
}

func handleEarlyAccessProposal(ctx sdk.Context, k keeper.Keeper, p *types.EarlyAccessProposal) error {
	for _, addr := range p.Users {
		user, err := k.GetUserFromString(ctx, addr)
		if err != nil {
			return err
		}

		keeper.AddEarlyAccessToUser(&user, "")

		k.SetUserFromUser(ctx, user)
	}

	return nil
}

type sortStruct struct {
	Id  uint64
	Set *types.Set
}
