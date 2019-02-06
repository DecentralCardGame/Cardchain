package cardservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type VoteRight struct {
	cardId string
	expireBlock int64
}

// NewHandler returns a handler for "cardservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgSaveCardContent:
			return handleMsgSaveCardContent(ctx, keeper, msg)
		//case MsgSetType:
			//return handleMsgSetType(ctx, keeper, msg)
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

/*
// Handle MsgSetType
func handleMsgSetType(ctx sdk.Context, keeper Keeper, msg MsgSetType) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.NameID)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetType(ctx, msg.NameID, msg.Value) // If so, set the type to the value specified in the msg.
	return sdk.Result{}                        // return
}
*/

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
	//stringId := strconv.FormatUint(currId, 10)

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

	newCard := Card{
		owner: msg.Buyer,
		content: []byte{},
		status: "scheme",
		votePool: sdk.NewInt64Coin("credits", 0),
		fairEnoughVotes: 0,
		overpoweredVotes: 0,
		underpoweredVotes: 0,
		nerflevel: 0,
	}

	keeper.SetCard(ctx, currId, newCard)

	return sdk.Result{}
}

func handleMsgSaveCardContent(ctx sdk.Context, keeper Keeper, msg MsgSaveCardContent) sdk.Result {
	card := keeper.GetCard(ctx, msg.CardId)

	if !msg.Owner.Equals(card.owner) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}



	return sdk.Result{}
}

// handle vote card message
func handleMsgVoteCard(ctx sdk.Context, keeper Keeper, msg MsgVoteCard) sdk.Result {

	// TODO activate vote right check, implement sdk.Errs
	/*
	voteRights := keeper.GetVoteRights(ctx, msg.Voter)
	rightsIndex := searchVoteRights(msg.CardId, voteRights)

	// check if voting rights are true
	if(rightsIndex < 0) {
		return sdk.ErrNoVoteRight("The right to vote on the card has expired")
	}

	//check if voting rights are timed out
	if(ctx.BlockHeight() > voteRights[rightsIndex].expireBlock) {
		return sdk.ErrVoteRightHasExpired("The right to vote on the card has expired")
	}
	*/
	/*
	// if the vote right is valid, get the Card
	keeper.GetCard(msg.CardID)

	// check for bounty
	if(msg.Card.votePool > 0) {
		msg.Voter add Coins from votepool
	}
	*/
	return sdk.Result{}
}
/*
func searchVoteRights(cardId string, rights []voteRight) int {
    for i, b := range rights {
        if b.cardId == cardId {
            return i
        }
    }
    return -1
}
*/
