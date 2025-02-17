package keeper

import (
	"context"
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CardSaveContent(goCtx context.Context, msg *types.MsgCardSaveContent) (*types.MsgCardSaveContentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	card := k.CardK.Get(ctx, msg.CardId)
	councilEnabled, err := k.FeatureFlagModuleInstance.Get(ctx, string(types.FeatureFlagName_Council))
	if err != nil {
		return nil, err
	}

	msgOwner, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	// Checks card status
	if councilEnabled && card.Status != types.CardStatus_prototype && card.Status != types.CardStatus_scheme {
		return nil, errorsmod.Wrap(types.ErrInvalidCardStatus, "Card has to be a prototype or scheme to be changeable")
	}

	if card.Owner == "" {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Card has no owner")
	} else if msg.Creator != card.Owner { // Checks if the the msg sender is the same as the current owner
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	cardobj, err := keywords.Unmarshal(msg.Content)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrCardobject, err.Error())

	}

	card.Content, err = json.Marshal(cardobj)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrCardobject, err.Error())
	}

	card.Notes = msg.Notes
	card.Artist = msg.Artist
	card.BalanceAnchor = msg.BalanceAnchor
	if card.Status == types.CardStatus_scheme {
		err = msgOwner.SchemeToCard(msg.CardId)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "An error accured while converting a card to a scheme: "+err.Error())
		}
	}

	claimedAirdrop := k.ClaimAirDrop(ctx, &msgOwner, types.AirDrop_create)
	k.SetUserFromUser(ctx, msgOwner)

	// Sets correct state
	if councilEnabled {
		card.Status = types.CardStatus_prototype
	} else {
		card.Status = types.CardStatus_permanent
	}
	k.CardK.Set(ctx, msg.CardId, card)
	k.SetLastCardModifiedNow(ctx)

	return &types.MsgCardSaveContentResponse{AirdropClaimed: claimedAirdrop}, nil
}
