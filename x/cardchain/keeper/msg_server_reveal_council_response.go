package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RevealCouncilResponse(goCtx context.Context, msg *types.MsgRevealCouncilResponse) (*types.MsgRevealCouncilResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collateralDeposit := k.GetParams(ctx).CollateralDeposit

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	council := k.GetCouncil(ctx, msg.CouncilId)
	if !stringItemInList(msg.Creator, council.Voters) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid Voter")
	}

	if council.Status != types.CouncelingStatus_commited {
		return nil, sdkerrors.Wrapf(types.ErrCouncilStatus, "Have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_commited.String())
	}

	hashStringResponse := GetResponseHash(msg.Response, msg.Secret)

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
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Already voted")
	}

	resp := types.WrapClearResponse{msg.Creator, msg.Response}
	council.ClearResponses = append(council.ClearResponses, &resp)

	if len(council.ClearResponses) == 5 {
		council.Status = types.CouncelingStatus_revealed
	}

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{collateralDeposit})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Voter does not have enough coins")
	}
	council.Treasury = council.Treasury.Add(collateralDeposit)

	//
	if len(council.ClearResponses) == 5 {
		var (
			nrNo, nrYes        int
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
		bounty := MulCoin(collateralDeposit, 2)
		votePool := MulCoin(collateralDeposit, 5)
		if nrNo > nrYes {
			for _, user := range deniers {
				err = k.MintCoinsToString(ctx, user, sdk.Coins{bounty})
				if err != nil {
					return nil, err
				}
				council.Treasury = council.Treasury.Sub(bounty)
			}
			k.AddPoolCredits(ctx, PublicPoolKey, council.Treasury)
			council.Treasury = council.Treasury.Sub(council.Treasury)
			council.Status = types.CouncelingStatus_councilClosed
		} else {
			card := k.GetCard(ctx, council.CardId)
			card.ResetVotes()
			card.VotePool = card.VotePool.Add(votePool)
			card.Status = types.Status_trial
			k.SetCard(ctx, council.CardId, card)
			council.TrialStart = uint64(ctx.BlockHeight())
			council.Treasury = council.Treasury.Sub(votePool)
		}
		for _, addr := range council.Voters {
			user, err := k.GetUserFromString(ctx, addr)
			if err != nil {
				return nil, err
			}
			user.CouncilStatus = types.CouncilStatus_available
			user.Cards = append(user.Cards, council.CardId)
			k.SetUserFromUser(ctx, user)
		}
	}

	k.SetCouncil(ctx, msg.CouncilId, council)

	return &types.MsgRevealCouncilResponseResponse{}, nil
}
