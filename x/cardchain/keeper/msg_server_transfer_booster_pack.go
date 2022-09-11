package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

func (k msgServer) TransferBoosterPack(goCtx context.Context, msg *types.MsgTransferBoosterPack) (*types.MsgTransferBoosterPackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	receiver, err := k.GetUserFromString(ctx, msg.Receiver)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if uint64(len(creator.BoosterPacks)) <= msg.BoosterPackId {
		return nil, sdkerrors.Wrapf(
			types.ErrBoosterPack,
			"BoosterPackId %d is not in list length %d",
			msg.BoosterPackId,
			len(creator.BoosterPacks),
		)
	}

	receiver.BoosterPacks = append(
		receiver.BoosterPacks,
		creator.BoosterPacks[msg.BoosterPackId],
	)

	creator.BoosterPacks = slices.Delete(
		creator.BoosterPacks,
		int(msg.BoosterPackId),
		int(msg.BoosterPackId+1),
	)

	k.SetUserFromUser(ctx, creator)
	k.SetUserFromUser(ctx, receiver)

	return &types.MsgTransferBoosterPackResponse{}, nil
}
