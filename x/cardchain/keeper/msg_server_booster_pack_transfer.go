package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BoosterPackTransfer(goCtx context.Context, msg *types.MsgBoosterPackTransfer) (*types.MsgBoosterPackTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	receiver, err := k.GetUserFromString(ctx, msg.Receiver)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	if uint64(len(creator.BoosterPacks)) <= msg.BoosterPackId {
		return nil, errorsmod.Wrapf(
			types.ErrInvalidData,
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

	return &types.MsgBoosterPackTransferResponse{}, nil
}
