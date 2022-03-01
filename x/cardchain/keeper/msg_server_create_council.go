package keeper

import (
	"context"
	"math/rand"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCouncil(goCtx context.Context, msg *types.MsgCreateCouncil) (*types.MsgCreateCouncilResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	card := k.GetCard(ctx, msg.CardId)
	if card.Status != types.Status_prototype {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCardStatus, "%s", card.Status.String())
	} else if card.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Creator")
	}

	var possibleVoters []User
	var fiveVoters []User
	var voters []string
	var status types.CouncelingStatus
	collateralDeposit := k.GetParams(ctx).CollateralDeposit
	treasury := MulCoin(collateralDeposit, 10)
	councilId := k.GetCouncilsNumber(ctx)
	users, addresses := k.GetAllUsers(ctx)

	for idx, user := range users {
		if user.CouncilStatus == types.CouncilStatus_available {
			possibleVoters = append(possibleVoters, User{*user, addresses[idx]})
		}
	}

	rand.Shuffle(len(possibleVoters), func(i, j int) {
		possibleVoters[i], possibleVoters[j] = possibleVoters[j], possibleVoters[i]
	})

	if len(possibleVoters) < 5 {
		status = types.CouncelingStatus_councilOpen
		fiveVoters = possibleVoters
	} else {
		status = types.CouncelingStatus_councilCreated
		fiveVoters = possibleVoters[:5]
	}

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{treasury})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Creator does not have enough coins")
	}

	for _, user := range fiveVoters {
		if status == types.CouncelingStatus_councilOpen {
			user.CouncilStatus = types.CouncilStatus_openCouncil
		} else {
			user.CouncilStatus = types.CouncilStatus_startedCouncil
		}
		k.SetUserFromUser(ctx, user)
		voters = append(voters, user.Addr.String())
	}

	council := types.Council{
		CardId:   msg.CardId,
		Voters:   voters,
		Treasury: treasury,
		Status:   status,
	}

	k.SetCouncil(ctx, councilId, council)

	return &types.MsgCreateCouncilResponse{}, nil
}
