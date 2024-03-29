package keeper

import (
	"context"
	"slices"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetProfileCard(goCtx context.Context, msg *types.MsgSetProfileCard) (*types.MsgSetProfileCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	card := k.Cards.Get(ctx, msg.CardId)
	if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "You have to own the card")
	} else if !slices.Contains([]types.Status{types.Status_trial, types.Status_permanent}, card.Status) {
		return nil, sdkerrors.Wrap(errors.ErrUnauthorized, "The card has to be playable")
	}

	user.ProfileCard = msg.CardId

	k.SetUserFromUser(ctx, user)

	return &types.MsgSetProfileCardResponse{}, nil
}
