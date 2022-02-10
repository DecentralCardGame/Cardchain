package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyCollection(goCtx context.Context, msg *types.MsgBuyCollection) (*types.MsgBuyCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	creator, err := k.GetUser(ctx, creatorAddr)
	if err != nil {
		return nil, err
	}

	collection, err := k.GetCurrentCollection(ctx)
	if err != nil {
		return nil, err
	}

	k.BurnCoinsFromAddr(ctx, creatorAddr, sdk.Coins{sdk.NewInt64Coin("credits", collectionPrice)})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Buyer does not have enough coins")
	}

	creator.Cards = append(creator.Cards, collection.Cards...)

	k.SetUser(ctx, creatorAddr, creator)

	return &types.MsgBuyCollectionResponse{}, nil
}