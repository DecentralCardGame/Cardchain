package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RevealCouncilResponse(goCtx context.Context, msg *types.MsgRevealCouncilResponse) (*types.MsgRevealCouncilResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	council := k.GetCouncil(ctx, msg.CouncilId)
	if !stringItemInList(msg.Creator, council.Voters) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid Voter")
	}

	response := msg.Response.String()
	clearResponse := response + msg.Secret
	hashResponse := sha256.Sum256([]byte(clearResponse))
	hashStringResponse := hex.EncodeToString(hashResponse[:])

	var origHash string
	for _, hash := range council.HashResponses {
		if hash.User == msg.Creator {
			origHash = hash.Hash
			break
		}
	}

	if origHash != hashStringResponse {
		return nil, types.ErrBadReveal
	}

	var allreadyVoted []string
	for _, response := range council.ClearResponses {
		allreadyVoted = append(allreadyVoted, response.User)
	}

	if stringItemInList(msg.Creator, allreadyVoted) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Allready voted")
	}

	resp := types.WrapClearResponse{msg.Creator, msg.Response}
	council.ClearResponses = append(council.ClearResponses, &resp)

	if len(council.ClearResponses) == 5 {
		council.Status = types.CouncelingStatus_revealed
	}
	k.SetCouncil(ctx, msg.CouncilId, council)

	return &types.MsgRevealCouncilResponseResponse{}, nil
}
