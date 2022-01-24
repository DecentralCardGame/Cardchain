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
		case *types.MsgVoteCard:
			return handleMsgVoteCard(ctx, k, msg)
		case *types.MsgSaveCardContent:
			return handleMsgSaveCardContent(ctx, k, msg)
			// this line is used by starport scaffolding # 1
		case *types.MsgCreateuser:
			return handleMsgCreateUser(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgSaveCardContent(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgSaveCardContent) (*sdk.Result, error) {
	card := keeper.GetCard(ctx, msg.CardId)
	msgOwner, _ := sdk.AccAddressFromBech32(msg.Owner)
	cardOwner, _ := sdk.AccAddressFromBech32(card.Owner)

	if !msgOwner.Equals(cardOwner) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	// TODO cards get a starting pool currently, this should be removed later and the starting pool should come after council decision
	card.VotePool.Add(sdk.NewInt64Coin("credits", 10))

	card.Content = []byte(msg.Content)
	card.Image = []byte(msg.Image)
	card.Status = "prototype"
	card.Notes = msg.Notes
	card.FullArt = msg.FullArt
	keeper.SetCard(ctx, msg.CardId, card)
	keeper.TransferSchemeToCard(ctx, msg.CardId, msgOwner)

	return &sdk.Result{}, nil
}

// handle vote card message
func handleMsgVoteCard(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgVoteCard) (*sdk.Result, error) {
	voter, _ := sdk.AccAddressFromBech32(msg.Voter)
	voteRights := keeper.GetVoteRights(ctx, voter)
	rightsIndex := SearchVoteRights(msg.CardId, voteRights)

	// check if voting rights are true
	if rightsIndex < 0 {
		return nil, sdkerrors.Wrap(types.ErrVoterHasNoVotingRights, "No Voting Rights")
	}

	//check if voting rights are timed out
	if ctx.BlockHeight() > voteRights[rightsIndex].ExpireBlock {
		return nil, sdkerrors.Wrap(types.ErrVoteRightIsExpired, "Voting Right has expired")
	}

	// if the vote right is valid, get the Card
	card := keeper.GetCard(ctx, msg.CardId)

	// check if card status is valid // TODO remove prototype as soon as the council exists
	if card.Status != "permanent" && card.Status != "trial" && card.Status != "prototype" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Voting on a card is only possible if it is in trial or a permanent card")
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	// check for specific bounty on the card
	if !card.VotePool.IsZero() {
		card.VotePool.Sub(sdk.NewInt64Coin("credits", 1))
		keeper.MintCoinsToAddr(ctx, voter, sdk.Coins{sdk.NewInt64Coin("credits", 1)})
	}

	// give generic bounty for voting
	keeper.MintCoinsToAddr(ctx, voter, sdk.Coins{sdk.NewInt64Coin("credits", 1)})

	keeper.SetCard(ctx, msg.CardId, card)

	keeper.RemoveVoteRight(ctx, voter, rightsIndex)

	return &sdk.Result{}, nil
}

func SearchVoteRights(cardID uint64, rights []*types.VoteRight) int {
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