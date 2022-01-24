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
		case *types.MsgBuyCardScheme:
			return handleMsgBuyCardScheme(ctx, k, msg)
			// this line is used by starport scaffolding # 1
		case *types.MsgCreateuser:
			return handleMsgCreateUser(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// Handle a message to buy name
func handleMsgBuyCardScheme(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgBuyCardScheme) (*sdk.Result, error) {
	lastId := keeper.GetLastCardSchemeId(ctx) // first get last card bought id
	currId := lastId + 1                      // iterate it by 1

	price := keeper.GetCardAuctionPrice(ctx)

	buyer, _ := sdk.AccAddressFromBech32(msg.Buyer)

	if msg.Bid.IsLT(price) { // Checks if the bid is less than price
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
	}

	err := keeper.BurnCoinsFromAddr(ctx, buyer, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Buyer does not have enough coins")
	}

	keeper.AddPublicPoolCredits(ctx, price)
	keeper.SetCardAuctionPrice(ctx, price.Add(price))
	keeper.SetLastCardSchemeId(ctx, currId)

	newCard := types.NewCard(buyer)

	keeper.SetCard(ctx, currId, newCard)
	keeper.AddOwnedCardScheme(ctx, currId, buyer)

	return &sdk.Result{}, nil
}

func handleMsgCreateUser(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgCreateuser) (*sdk.Result, error) {
	// check if user already exists
	keeper.Createuser(ctx, msg.GetNewUserAddr(), msg.Alias)

	// this has been moved to keeper.InitUser, but maybe will be back here some day?
	// give starting credits
	//if(!keeper.GetPublicPoolCredits(ctx).IsZero()) {
	//keeper.SubtractPublicPoolCredits(ctx, sdk.NewInt64Coin("credits", 1))
	//keeper.coinKeeper.AddCoins(ctx, msg.NewUser, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	//}

	return &sdk.Result{}, nil
}
