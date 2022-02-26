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

	colDep := k.GetParams(ctx).CollateralDeposit

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

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

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{colDep})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Voter does not have enough coins")
	}
	council.Treasury = council.Treasury.Add(colDep)

	//
	if len(council.ClearResponses) == 5 {
		var (
			nrNo, nrYes int
			approvers, deniers []string
		)
		for _, response := range council.ClearResponses {
			if response.Response == types.Response_Yes {
				nrYes++
				approvers = append(approvers, response.User)
			} else {
				nrNo++
				deniers = append(deniers, response.User)
			}
		}
		bounty := sdk.Coin{colDep.Denom, colDep.Amount.Mul(sdk.NewInt(2))}
		votePool := sdk.Coin{colDep.Denom, colDep.Amount.Mul(sdk.NewInt(5))}
		if nrNo > nrYes {
			for _, user := range deniers {
				addr, err := sdk.AccAddressFromBech32(user)
				if err != nil {
					return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
				}
				k.MintCoinsToAddr(ctx, addr, sdk.Coins{bounty})
				council.Treasury = council.Treasury.Sub(bounty)
			}
			pool := k.GetPool(ctx, PublicPoolKey)
			pool = pool.Add(council.Treasury)
			k.SetPool(ctx, PublicPoolKey, pool)
			council.Treasury = council.Treasury.Sub(council.Treasury)
			council.Status = types.CouncelingStatus_councilClosed
		} else {
			card := k.GetCard(ctx, council.CardId)
			card.VotePool = card.VotePool.Add(votePool)
			card.Status = types.Status_trial
			k.SetCard(ctx, council.CardId, card)
			council.Treasury = council.Treasury.Sub(votePool)
		}
	}

	k.SetCouncil(ctx, msg.CouncilId, council)

	return &types.MsgRevealCouncilResponseResponse{}, nil
}
