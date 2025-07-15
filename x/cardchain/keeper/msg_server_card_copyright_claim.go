package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CardCopyrightClaim(goCtx context.Context, msg *types.MsgCardCopyrightClaim) (*types.MsgCardCopyrightClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "expected %s got %s", k.authority, msg.Authority)
	}

	card := k.CardK.Get(ctx, msg.CardId)
	if card.Status == types.CardStatus_none {
		return nil, errorsmod.Wrapf(types.ErrCardDoesNotExist, "cardstatus is none")
	}

	image := k.Images.Get(ctx, card.ImageId)

	image.Image = []byte{}
	card.Artist = card.Owner
	card.Status = types.CardStatus_suspended

	k.SetLastCardModifiedNow(ctx)
	k.CardK.Set(ctx, msg.CardId, card)
	k.Images.Set(ctx, card.ImageId, image)

	return &types.MsgCardCopyrightClaimResponse{}, nil
}
