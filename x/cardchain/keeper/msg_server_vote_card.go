package keeper

import (
	"context"
	"fmt"

	"golang.org/x/exp/slices"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) VoteCard(goCtx context.Context, msg *types.MsgVoteCard) (*types.MsgVoteCardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	voter, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	voteRights := k.GetVoteRights(ctx, voter)
	rightsIndex := slices.IndexFunc(voteRights, func(s *types.VoteRight) bool {return s.CardId == msg.CardId})

	// check if voting rights are true
	if rightsIndex < 0 {
		return nil, sdkerrors.Wrap(types.ErrVoterHasNoVotingRights, "No Voting Rights")
	}

	//check if voting rights are timed out
	if ctx.BlockHeight() > voteRights[rightsIndex].ExpireBlock {
		return nil, sdkerrors.Wrap(types.ErrVoteRightIsExpired, "Voting Right has expired")
	}

	// if the vote right is valid, get the Card
	card := k.Card.Get(ctx, msg.CardId)

	// check if card status is valid
	if card.Status != types.Status_permanent && card.Status != types.Status_trial {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Voting on a card is only possible if it is in trial or a permanent card")
	}

	switch msg.VoteType {
	case "fair_enough":
		card.FairEnoughVotes++
	case "inappropriate":
		card.InappropriateVotes++
	case "overpowered":
		card.OverpoweredVotes++
	case "underpowered":
		card.UnderpoweredVotes++
	default:
		errMsg := fmt.Sprintf("Unrecognized card vote type: %s", msg.VoteType)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}
	card.Voters = append(card.Voters, msg.Creator)

	// check for specific bounty on the card
	if !card.VotePool.IsZero() {
		err := k.MintCoinsToAddr(ctx, voter, sdk.Coins{sdk.NewInt64Coin("ucredits", 1000000)})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
		}
		card.VotePool.Sub(sdk.NewInt64Coin("ucredits", 1000000))
	}

	amount := k.GetVoteReward(ctx)
	k.MintCoinsToAddr(ctx, voter, sdk.Coins{amount})
	k.SubPoolCredits(ctx, BalancersPoolKey, amount)

	k.Card.Set(ctx, msg.CardId, card)

	votes := k.GetRunningAverage(ctx, Votes24ValueKey)
	votes.Arr[len(votes.Arr)-1]++
	k.SetRunningAverage(ctx, Votes24ValueKey, votes)

	err = k.RemoveVoteRight(ctx, voter, rightsIndex)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUserDoesNotExist, err.Error())
	}

	return &types.MsgVoteCardResponse{}, nil
}
