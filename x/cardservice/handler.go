package cardservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
		case MsgTransferCard:
			return handleMsgTransferCard(ctx, keeper, msg)
		case MsgDonateToCard:
			return handleMsgDonateToCard(ctx, keeper, msg)
		case MsgCreateUser:
			return handleMsgCreateUser(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized cardservice Msg type in handler.go: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle MsgBuyName
func handleMsgBuyCardScheme(ctx sdk.Context, keeper Keeper, msg MsgBuyCardScheme) sdk.Result {

	lastId := keeper.GetLastCardSchemeId(ctx) // first get last card bought id
	currId := lastId + 1                      // iterate it by 1
	//stringId := strconv.FormatUint(currId, 10)

	price := keeper.GetCardAuctionPrice(ctx)

	if price.IsGTE(msg.Bid) { // Checks if the the bid price is greater than the price paid by the current owner
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}

	_, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
	}

	keeper.AddPublicPoolCredits(ctx, price)
	keeper.SetCardAuctionPrice(ctx, price.Add(price))
	keeper.SetLastCardSchemeId(ctx, currId)

	newCard := NewCard(msg.Buyer)

	keeper.SetCard(ctx, currId, newCard)
	keeper.AddOwnedCardScheme(ctx, currId, msg.Buyer)

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

	card.Content = []byte(msg.Content)
	card.Status = "prototype"
	keeper.SetCard(ctx, msg.CardId, card)
	keeper.TransferSchemeToCard(ctx, msg.CardId, msg.Owner)

	return sdk.Result{}
}

// handle vote card message
func handleMsgVoteCard(ctx sdk.Context, keeper Keeper, msg MsgVoteCard) sdk.Result {
	voteRights := keeper.GetVoteRights(ctx, msg.Voter)
	rightsIndex := SearchVoteRights(msg.CardId, voteRights)

	// check if voting rights are true
	if rightsIndex < 0 {
		return ErrNoVotingRights().Result()
	}

	//check if voting rights are timed out
	if ctx.BlockHeight() > voteRights[rightsIndex].ExpireBlock {
		return ErrVoteRightHasExpired().Result()
	}

	// if the vote right is valid, get the Card
	card := keeper.GetCard(ctx, msg.CardId)

	// check if card status is valid // TODO remove prototype as soon as the council exists
	if card.Status != "permanent" && card.Status != "trial" && card.Status != "prototype" {
		return sdk.ErrUnknownRequest("Voting on a card is only possible if it is in trial or a permanent card").Result()
	}

	switch msg.VoteType {
	case "fair_enough":
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
	if !card.VotePool.IsZero() {
		card.VotePool.Sub(sdk.NewInt64Coin("credits", 1))
		keeper.coinKeeper.AddCoins(ctx, msg.Voter, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	}

	keeper.SetCard(ctx, msg.CardId, card)

	keeper.RemoveVoteRight(ctx, msg.Voter, rightsIndex)

	return sdk.Result{}
}

func SearchVoteRights(cardID uint64, rights []VoteRight) int {
	if rights == nil {
		return -1
	}
	for i, b := range rights {
		if b.CardId == cardID {
			return i
		}
	}
	return -1
}

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
	card.VotePool = card.VotePool.Add(msg.Amount)
	keeper.SetCard(ctx, msg.CardId, card)

	return sdk.Result{}
}

// handle create user message
func handleMsgCreateUser(ctx sdk.Context, keeper Keeper, msg MsgCreateUser) sdk.Result {

	// check if Creator is valid, later this is for preventing sybil attack
	if true {

	}

	// check if user already exists
	keeper.CreateUser(ctx, msg.NewUser, msg.Alias)

	// this has been moved to keeper.InitUser, but maybe will be back here some day?
	// give starting credits
	//if(!keeper.GetPublicPoolCredits(ctx).IsZero()) {
	//keeper.SubtractPublicPoolCredits(ctx, sdk.NewInt64Coin("credits", 1))
	//keeper.coinKeeper.AddCoins(ctx, msg.NewUser, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	//}

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
