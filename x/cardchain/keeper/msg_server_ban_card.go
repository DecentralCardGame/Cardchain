package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BanCard(goCtx context.Context, msg *types.MsgBanCard) (*types.MsgBanCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "expected %s got %s", k.authority, msg.Authority)
	}

	card := k.CardK.Get(ctx, msg.CardId)

	if card.Status == types.CardStatus_none {
		return nil, errorsmod.Wrapf(types.ErrCardDoesNotExist, "cardstatus is none")
	}

	card.Status = types.CardStatus_banned

	k.CardK.Set(ctx, msg.CardId, card)

	return &types.MsgBanCardResponse{}, nil
}
