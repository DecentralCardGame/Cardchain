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
			case MsgSaveCardContent:
				return handleMsgSaveCardContent(ctx, keeper, msg)
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
	keeper.SetCardAuctionPrice(ctx, price.Add(price))
	keeper.SetLastCardSchemeId(ctx, currId)

	newCard := NewCard(msg.Buyer)

	keeper.SetCard(ctx, currId, newCard)

	return sdk.Result{}
}

func handleMsgSaveCardContent(ctx sdk.Context, keeper Keeper, msg MsgSaveCardContent) sdk.Result {

	card := keeper.GetCard(ctx, msg.CardId)

	if !msg.Owner.Equals(card.Owner) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}

	// TODO card content should be deserialized and serialized here
	// serialize it
	// ...

	// TODO cards get a starting pool currently, this should be removed later and the starting pool should come after council decision
	card.VotePool.Add(sdk.NewInt64Coin("credits", 10))

	card.Content = msg.Content
	keeper.SetCard(ctx, msg.CardId, card)

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

	// if the vote right is valid, get the Card
	card := keeper.GetCard(ctx, msg.CardId)

	// check if card status is valid // TODO enable
	/*
	if(card.Status != "permanent" && card.Status != "trial") {
		return sdk.ErrUnknownRequest("Voting on a card is only possible if it is in trial or a permanent card").Result()
	}
	*/

	switch msg.VoteType {
	case "fairenough":
		card.FairEnoughVotes++
	case "inappropriate":
		card.InappropriateVotes++
	case "overpowered":
		card.OverpoweredVotes++
	case "underpowered":
		card.UnderpoweredVotes++
	default:
		errMsg := fmt.Sprintf("Unrecognized card vote type: %s", msg.VoteType)
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	// check for bounty
	if(!card.VotePool.IsZero()) {
		card.VotePool.Sub(sdk.NewInt64Coin("credits", 1))
		keeper.coinKeeper.AddCoins(ctx, msg.Voter, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	}

	keeper.SetCard(ctx, msg.CardId, card)

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

// handle vote card message
func handleMsgTransferCard(ctx sdk.Context, keeper Keeper, msg MsgTransferCard) sdk.Result {

	// if the vote right is valid, get the Card
	card := keeper.GetCard(ctx, msg.CardId)

	// check if card status is valid // TODO ponder if this is necessary for transfer card
	/*
	if(card.Status != "permanent" && card.Status != "trial") {
		return sdk.ErrUnknownRequest("Transferring a card is only possible if it is in trial or a permanent card").Result()
	}
	*/

	if !msg.Sender.Equals(card.Owner) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}

	card.Owner = msg.Receiver
	keeper.SetCard(ctx, msg.CardId, card)

	return sdk.Result{}
}

// handle donate to card message
func handleMsgDonateToCard(ctx sdk.Context, keeper Keeper, msg MsgDonateToCard) sdk.Result {

	_, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Donator, sdk.Coins{msg.Amount}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInsufficientCoins("Donator does not have enough coins").Result()
	}

	card := keeper.GetCard(ctx, msg.CardId)
	card.VotePool.Add(msg.Amount)
	keeper.SetCard(ctx, msg.CardId, card)

	return sdk.Result{}
}

// handle create user message
func handleMsgCreateUser(ctx sdk.Context, keeper Keeper, msg MsgCreateUser) sdk.Result {
	newUser := NewUser()
	keeper.SetUser(ctx, msg.NewUser, newUser)

	// give starting credits
	if(!keeper.GetPublicPoolCredits(ctx).IsZero()) {
		keeper.DeltaPublicPoolCredits(ctx, sdk.NewInt64Coin("credits", -1))
		keeper.coinKeeper.AddCoins(ctx, msg.NewUser, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	}

	return sdk.Result{}
}

/*
func sellCardInstances(ctx sdk.Context, keeper Keeper, msg MsgSellCardInstances) sdk.Result {

	_, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, stringId), msg.Bid)
	if err != nil {
		return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
	}

	return sdk.Result{}
}
*/
