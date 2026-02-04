package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const PICK_SCALER = 1.0

func (k msgServer) UpgradePicksReport(goCtx context.Context, msg *types.MsgUpgradePicksReport) (*types.MsgUpgradePicksReportResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrUserDoesNotExist, err.Error())
	}
	if creator.ReportMatches == false {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Reporter")
	}

	pickAddition := PICK_SCALER * float32(len(msg.Dismissed)) / float32(len(msg.Picked))

	for _, dismissed := range msg.Dismissed {
		upgradeFactor := k.UpgradeFactorK.Get(ctx, dismissed)
		upgradeFactor.Picks--
		k.UpgradeFactorK.Set(ctx, dismissed, upgradeFactor)
	}

	for _, picked := range msg.Picked {
		upgradeFactor := k.UpgradeFactorK.Get(ctx, picked)
		upgradeFactor.Picks += pickAddition
		k.UpgradeFactorK.Set(ctx, picked, upgradeFactor)
	}

	return &types.MsgUpgradePicksReportResponse{}, nil
}
