package cardservice

import (
	"fmt"
	"strconv"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "cardservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgSetType:
			return handleMsgSetType(ctx, keeper, msg)
		case MsgBuyCardScheme:
			return handleMsgBuyCardScheme(ctx, keeper, msg)
		case MsgVoteCard:
			return handleMsgVoteCard(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized cardservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle MsgSetType
func handleMsgSetType(ctx sdk.Context, keeper Keeper, msg MsgSetType) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.NameID)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetType(ctx, msg.NameID, msg.Value) // If so, set the type to the value specified in the msg.
	return sdk.Result{}                        // return
}

// Handle MsgSetName
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.NameID)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetName(ctx, msg.NameID, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                        // return
}

// Handle MsgBuyName
func handleMsgBuyCardScheme(ctx sdk.Context, keeper Keeper, msg MsgBuyCardScheme) sdk.Result {

	lastId := keeper.GetLastCardSchemeId(ctx) // first get last card bought id
	currId := lastId+1		// iterate it by 1
	stringId := strconv.FormatUint(currId, 10)

	price := keeper.GetCardAuctionPrice(ctx)

	if price.IsGTE(msg.Bid) { // Checks if the the bid price is greater than the price paid by the current owner
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	/*if keeper.HasOwner(ctx, stringId) {
		_, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, stringId), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {*/
		_, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	//}

	keeper.DeltaPublicPoolCredits(ctx, price)
	keeper.SetCardAuctionPrice(ctx, price.Plus(price))
	keeper.SetLastCardSchemeId(ctx, currId)
	keeper.SetOwner(ctx, stringId, msg.Buyer)
	//keeper.SetPrice(ctx, stringId, msg.Bid)
	return sdk.Result{}
}

// handle vote card message
func handleMsgVoteCard(ctx sdk.Context, keeper Keeper, msg MsgVoteCard) sdk.Result {

	voteRights := keeper.GetVoteRights(ctx, msg.CardId, msg.Voter)

	// check if voting rights are true
	if(msg.Voter.voteRights contains msg.CardId) {
		//check if voting rights are timed out
		if(ctx.BlockHeight() - msg.Card.voteRights.blockheight < params.cardVoteDuration) {

		}
	}

	// check for bounty
	if(msg.Card.votePool > 0) {
		msg.Voter add Coins from votepool
	}

	sdk.Result{}
}

func VoterInVoteRights(cardId string, rights []voteRights) bool {
    for _, b := range rights {
        if b. == cardId {
            return true
        }
    }
    return false
}
