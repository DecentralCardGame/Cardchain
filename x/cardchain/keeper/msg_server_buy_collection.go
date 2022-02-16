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

	// contribs := k.GetAllCollectionContributors(ctx, collection)
	//
	// for _, contrib := range contribs {
	// 	contribAddr, err := sdk.AccAddressFromBech32(contrib)
	// 	if err != nil {
	// 		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	// 	}
	//
	// 	err = k.BankKeeper.SendCoins(ctx, creator.Addr, contribAddr, sdk.DecCoins{sdk.NewDecCoinFromDec("credits", sdk.NewDecWithPrec(int64(k.GetParams(ctx).CollectionPrice*1000/int64(len(contribs))), 3))})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{sdk.NewInt64Coin("credits", k.GetParams(ctx).CollectionPrice)}) // TODO Profitdividing
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Buyer does not have enough coins")
	}

	creator.Cards = append(creator.Cards, collection.Cards...)

	k.SetUserFromUser(ctx, creator)

	return &types.MsgBuyCollectionResponse{}, nil
}
