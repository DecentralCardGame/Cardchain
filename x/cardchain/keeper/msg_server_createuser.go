package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Createuser(goCtx context.Context, msg *types.MsgCreateuser) (*types.MsgCreateuserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if user already exists
	k.CreateUser(ctx, msg.GetNewUserAddr(), msg.Alias)

	// this has been moved to keeper.InitUser, but maybe will be back here some day?
	// give starting credits
	//if(!keeper.GetPublicPoolCredits(ctx).IsZero()) {
	//keeper.SubtractPublicPoolCredits(ctx, sdk.NewInt64Coin("credits", 1))
	//keeper.coinKeeper.AddCoins(ctx, msg.NewUser, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	//}

	return &types.MsgCreateuserResponse{}, nil
}
