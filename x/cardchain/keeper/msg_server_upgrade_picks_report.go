package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
)

func (k msgServer) UpgradePicksReport(goCtx context.Context, msg *types.MsgUpgradePicksReport) (*types.MsgUpgradePicksReportResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgUpgradePicksReportResponse{}, nil
}
