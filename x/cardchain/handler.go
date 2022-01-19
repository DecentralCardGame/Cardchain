package cardchain

import (
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	//msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// case *types.MsgCreateuser:
		// 	res, err := msgServer.Createuser(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		// 	// this line is used by starport scaffolding # 1
		case *types.MsgCreateuser:
			return handleMsgCreateUser(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}


func handleMsgCreateUser(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgCreateuser) (*sdk.Result, error) {
	// check if user already exists
	keeper.Createuser(ctx, msg.NewUser, msg.Alias)

	// this has been moved to keeper.InitUser, but maybe will be back here some day?
	// give starting credits
	//if(!keeper.GetPublicPoolCredits(ctx).IsZero()) {
	//keeper.SubtractPublicPoolCredits(ctx, sdk.NewInt64Coin("credits", 1))
	//keeper.coinKeeper.AddCoins(ctx, msg.NewUser, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	//}

	return &sdk.Result{}, nil
}
