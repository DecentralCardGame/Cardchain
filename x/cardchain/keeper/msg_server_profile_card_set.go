package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ProfileCardSet(goCtx context.Context, msg *types.MsgProfileCardSet) (*types.MsgProfileCardSetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	card := k.Cards.Get(ctx, msg.CardId)
	if card.Owner != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "You have to own the card")
	} else if !slices.Contains([]types.CardStatus{types.CardStatus_trial, types.CardStatus_permanent}, card.Status) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "The card has to be playable")
	}

	user.ProfileCard = msg.CardId

	k.SetUserFromUser(ctx, user)

	return &types.MsgProfileCardSetResponse{}, nil
}
