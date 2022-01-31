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
		// case *types.MsgCreateuser:  // Default handeling
		// 	res, err := msgServer.Createuser(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddArtwork:
			return handleMsgAddArtwork(ctx, k, msg)
		case *types.MsgBuyCardScheme:
			return handleMsgBuyCardScheme(ctx, k, msg)
		case *types.MsgVoteCard:
			return handleMsgVoteCard(ctx, k, msg)
		case *types.MsgSaveCardContent:
			return handleMsgSaveCardContent(ctx, k, msg)
		case *types.MsgTransferCard:
			return handleMsgTransferCard(ctx, k, msg)
		case *types.MsgDonateToCard:
			return handleMsgDonateToCard(ctx, k, msg)
		case *types.MsgCreateuser:
			return handleMsgCreateUser(ctx, k, msg)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgAddArtwork(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgAddArtwork) (*sdk.Result, error) {
	card := keeper.GetCard(ctx, msg.CardId)

	if card.Artist != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Artist")
	}

	card.FullArt = msg.FullArt
	card.Image = msg.Image

	keeper.SetCard(ctx, msg.CardId, card)

	return &sdk.Result{}, nil
}

func handleMsgDonateToCard(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgDonateToCard) (*sdk.Result, error) {
	donator, err := sdk.AccAddressFromBech32(msg.Donator)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	err = keeper.BurnCoinsFromAddr(ctx, donator, sdk.Coins{msg.Amount}) // If so, deduct the Bid amount from the sender
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Donator does not have enough coins")
	}

	card := keeper.GetCard(ctx, msg.CardId)
	card.VotePool = card.VotePool.Add(msg.Amount)
	keeper.SetCard(ctx, msg.CardId, card)

	return &sdk.Result{}, nil
}

func handleMsgTransferCard(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgTransferCard) (*sdk.Result, error) {

	// if the vote right is valid, get the Card
	card := keeper.GetCard(ctx, msg.CardId)

	// check if card status is valid // TODO ponder if this is necessary for transfer card
	/*
		if(card.Status != "permanent" && card.Status != "trial") {
			return sdk.ErrUnknownRequest("Transferring a card is only possible if it is in trial or a permanent card").Result()
		}
	*/
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	owner, err := sdk.AccAddressFromBech32(card.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if !sender.Equals(owner) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	card.Owner = msg.Receiver
	keeper.SetCard(ctx, msg.CardId, card)

	return &sdk.Result{}, nil
}

func handleMsgSaveCardContent(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgSaveCardContent) (*sdk.Result, error) {
	card := keeper.GetCard(ctx, msg.CardId)

	msgOwner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, err.Error())
	}

	// TODO Add error when writing to unowned card

	cardOwner, err := sdk.AccAddressFromBech32(card.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, err.Error())
	}

	if !msgOwner.Equals(cardOwner) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	// TODO cards get a starting pool currently, this should be removed later and the starting pool should come after council decision
	card.VotePool.Add(sdk.NewInt64Coin("credits", 10))

	card.Content = []byte(msg.Content)
	// card.Image = []byte(msg.Image)
	card.Status = "prototype"
	card.Notes = msg.Notes
	card.Artist = msg.Artist
	// card.FullArt = msg.FullArt
	keeper.SetCard(ctx, msg.CardId, card)
	keeper.TransferSchemeToCard(ctx, msg.CardId, msgOwner)

	return &sdk.Result{}, nil
}

// handle vote card message
func handleMsgVoteCard(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgVoteCard) (*sdk.Result, error) {
	voter, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

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

	err = keeper.RemoveVoteRight(ctx, voter, rightsIndex)
	if err != nil {
		return nil, err
	}

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

	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}

	if msg.Bid.IsLT(price) { // Checks if the bid is less than price
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
	}

	err = keeper.BurnCoinsFromAddr(ctx, buyer, sdk.Coins{price}) // If so, deduct the Bid amount from the sender
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

// apply nerf levels and remove inappropriate cards
func UpdateNerfLevels(ctx sdk.Context, keeper keeper.Keeper) sdk.Result {

	buffbois, nerfbois, _, banbois := keeper.GetOPandUPCards(ctx)
	/*
		fmt.Println("buff:")
		fmt.Println(buffbois)
		fmt.Println("nerf:")
		fmt.Println(nerfbois)
		fmt.Println("fair:")
		fmt.Println(fairbois)
		fmt.Println("ban:")
		fmt.Println(banbois)
	*/
	keeper.NerfBuffCards(ctx, buffbois, true)
	keeper.NerfBuffCards(ctx, nerfbois, false)
	keeper.UpdateBanStatus(ctx, banbois)

	keeper.ResetAllVotes(ctx)

	return sdk.Result{}
}
