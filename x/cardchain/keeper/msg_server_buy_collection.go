package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCollection(goCtx context.Context, msg *types.MsgBuyCollection) (*types.MsgBuyCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	collection := k.GetCollection(ctx, msg.CollectionId)

	if collection.Status != types.CStatus_active {
		return nil, types.ErrNoActiveCollection
	}

	k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{sdk.NewInt64Coin("credits", k.GetParams(ctx).CollectionPrice)}) // TODO Profitdividing
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Buyer does not have enough coins")
	}

	creator.Cards = append(creator.Cards, collection.Cards...)

	k.SetUserFromUser(ctx, creator)

	return &types.MsgBuyCollectionResponse{}, nil
}
