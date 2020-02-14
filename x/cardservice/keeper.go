package cardservice

import (
	//"fmt"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	cardsStoreKey    sdk.StoreKey // Unexposed key to access card store from sdk.Context
	usersStoreKey    sdk.StoreKey // Unexposed key to access user store from sdk.Context
	internalStoreKey sdk.StoreKey // Unexposed key to access internal variables from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the cardservice Keeper
func NewKeeper(coinKeeper bank.Keeper, cardsStoreKey sdk.StoreKey, usersStoreKey sdk.StoreKey, internalStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:       coinKeeper,
		cardsStoreKey:    cardsStoreKey,
		usersStoreKey:    usersStoreKey,
		internalStoreKey: internalStoreKey,
		cdc:              cdc,
	}
}

// GetLastCardScheme - get the current id of the last bought card scheme
func (k Keeper) GetLastCardSchemeId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("lastCardScheme"))
	return binary.BigEndian.Uint64(bz)
}

// SetLastCardScheme - sets the current id of the last bought card scheme
func (k Keeper) SetLastCardSchemeId(ctx sdk.Context, lastId uint64) {
	store := ctx.KVStore(k.internalStoreKey)
	store.Set([]byte("lastCardScheme"), sdk.Uint64ToBigEndian(lastId))
}

// returns the current price of the card scheme auction
func (k Keeper) GetCardAuctionPrice(ctx sdk.Context) sdk.Coin {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
	var price sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &price)
	return price
}

// sets the current price of the card scheme auction
func (k Keeper) SetCardAuctionPrice(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.internalStoreKey)
	store.Set([]byte("currentCardSchemeAuctionPrice"), k.cdc.MustMarshalBinaryBare(price))
}

// gets the credits in the public pool
func (k Keeper) GetPublicPoolCredits(ctx sdk.Context) sdk.Coin {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	return amount
}

// sets the credits in the public pool
func (k Keeper) SetPublicPoolCredits(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.internalStoreKey)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(price))
}

// subtracts credits from the public pool
func (k Keeper) SubtractPublicPoolCredits(ctx sdk.Context, delta sdk.Coin) {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	newAmount := amount.Sub(delta)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(newAmount))
}

// adds or subtracts credits from the public pool
func (k Keeper) AddPublicPoolCredits(ctx sdk.Context, delta sdk.Coin) {
	store := ctx.KVStore(k.internalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	newAmount := amount.Add(delta)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(newAmount))
}

func (k Keeper) GetCard(ctx sdk.Context, cardId uint64) Card {
	store := ctx.KVStore(k.cardsStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(cardId))

	var gottenCard Card
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenCard)
	return gottenCard
}

func (k Keeper) SetCard(ctx sdk.Context, cardId uint64, newCard Card) {
	store := ctx.KVStore(k.cardsStoreKey)
	store.Set(sdk.Uint64ToBigEndian(cardId), k.cdc.MustMarshalBinaryBare(newCard))
}

func (k Keeper) GetUser(ctx sdk.Context, address sdk.AccAddress) User {
	store := ctx.KVStore(k.usersStoreKey)
	bz := store.Get(address)

	var gottenUser User
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenUser)
	return gottenUser
}

func (k Keeper) AddVoteRightsToAllUsers(ctx sdk.Context, expireBlock int64) {
	votingRights := k.GetVoteRightToAllCards(ctx, expireBlock)

	userStore := ctx.KVStore(k.usersStoreKey)

	userIterator := sdk.KVStorePrefixIterator(userStore, nil)

	for ; userIterator.Valid(); userIterator.Next() {
		var user User
		k.cdc.MustUnmarshalBinaryBare(userIterator.Value(), &user)
		user.VoteRights = votingRights
		userStore.Set(userIterator.Key(), k.cdc.MustMarshalBinaryBare(user))
	}

	userIterator.Close()
}

func (k Keeper) GetVoteRightToAllCards(ctx sdk.Context, expireBlock int64) []VoteRight {
	cardStore := ctx.KVStore(k.cardsStoreKey)
	cardIterator := sdk.KVStorePrefixIterator(cardStore, nil)

	votingRights := []VoteRight{}

	for ; cardIterator.Valid(); cardIterator.Next() {
		right := NewVoteRight(binary.BigEndian.Uint64(cardIterator.Key()), expireBlock)
		votingRights = append(votingRights, right)
	}
	cardIterator.Close()

	return votingRights
}

func (k Keeper) RemoveVoteRight(ctx sdk.Context, userAddress sdk.AccAddress, rightsIndex int) {
	user := k.GetUser(ctx, userAddress)
	user.VoteRights[rightsIndex] = user.VoteRights[len(user.VoteRights)-1]
	//user.VoteRights[len(user.VoteRights)-1] = null
	user.VoteRights = user.VoteRights[:len(user.VoteRights)-1]
	userStore := ctx.KVStore(k.usersStoreKey)
	userStore.Set(userAddress, k.cdc.MustMarshalBinaryBare(user))
}

func (k Keeper) SetUser(ctx sdk.Context, address sdk.AccAddress, userData User) {
	store := ctx.KVStore(k.usersStoreKey)
	store.Set(address, k.cdc.MustMarshalBinaryBare(userData))
}

func (k Keeper) SetUserName(ctx sdk.Context, address sdk.AccAddress, name string) {
	store := ctx.KVStore(k.usersStoreKey)

	gottenUser := k.GetUser(ctx, address)

	gottenUser.Alias = name

	store.Set(address, k.cdc.MustMarshalBinaryBare(gottenUser))
}

func (k Keeper) InitUser(ctx sdk.Context, address sdk.AccAddress, alias string) {
	store := ctx.KVStore(k.usersStoreKey)

	newUser := NewUser()
	newUser.Alias = alias
	k.coinKeeper.AddCoins(ctx, address, sdk.Coins{sdk.NewInt64Coin("credits", 1000)})
	const votingRightsExpirationTime = 500
	newUser.VoteRights = k.GetVoteRightToAllCards(ctx, ctx.BlockHeight()+votingRightsExpirationTime)

	store.Set(address, k.cdc.MustMarshalBinaryBare(newUser))
}

func (k Keeper) CreateUser(ctx sdk.Context, newUser sdk.AccAddress, alias string) User {
	// check if user already exists
	store := ctx.KVStore(k.usersStoreKey)
	bz := store.Get(newUser)

	if bz == nil {
		k.InitUser(ctx, newUser, alias)
	}

	return k.GetUser(ctx, newUser)
}

func (k Keeper) AddOwnedCardScheme(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	store := ctx.KVStore(k.usersStoreKey)
	bz := store.Get(address)

	var gottenUser User
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenUser)

	gottenUser.OwnedCardSchemes = append(gottenUser.OwnedCardSchemes, cardId)

	store.Set(address, k.cdc.MustMarshalBinaryBare(gottenUser))
}

func (k Keeper) TransferSchemeToCard(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	store := ctx.KVStore(k.usersStoreKey)
	bz := store.Get(address)

	var gottenUser User
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenUser)

	idPosition := indexOfId(cardId, gottenUser.OwnedCardSchemes)

	if idPosition >= 0 {
			gottenUser.OwnedCards = append(gottenUser.OwnedCards, cardId)
			gottenUser.OwnedCardSchemes = append(gottenUser.OwnedCardSchemes[:idPosition], gottenUser.OwnedCardSchemes[idPosition+1:]...)

			store.Set(address, k.cdc.MustMarshalBinaryBare(gottenUser))
	}
}

func indexOfId(cardID uint64, cards []uint64) int {
	if cards == nil {
		return -1
	}
	for i, b := range cards {
		if b == cardID {
			return i
		}
	}
	return -1
}

func (k Keeper) GetVoteRights(ctx sdk.Context, voter sdk.AccAddress) []VoteRight {
	user := k.GetUser(ctx, voter)
	return user.VoteRights
}

func (k Keeper) SetVoteRights(ctx sdk.Context, voteRights []VoteRight, voter sdk.AccAddress) {
	store := ctx.KVStore(k.usersStoreKey)
	store.Set(voter, k.cdc.MustMarshalBinaryBare(voteRights))
}

func (k Keeper) NerfBuffCards(ctx sdk.Context, cardIds []uint64, buff bool) {
	store := ctx.KVStore(k.cardsStoreKey)

	for _, val := range cardIds {
		bz := store.Get(sdk.Uint64ToBigEndian(val))
		var buffCard Card
		k.cdc.MustUnmarshalBinaryBare(bz, &buffCard)

		if buff {
			buffCard.Nerflevel -= 1
		} else {
			buffCard.Nerflevel += 1
		}

		store.Set(sdk.Uint64ToBigEndian(val), k.cdc.MustMarshalBinaryBare(buffCard))
	}
}

func (k Keeper) ResetAllVotes(ctx sdk.Context) {
	store := ctx.KVStore(k.cardsStoreKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)

	for ; iterator.Valid(); iterator.Next() {
		var resetCard Card
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &resetCard)

		resetCard.FairEnoughVotes = 0
		resetCard.OverpoweredVotes = 0
		resetCard.UnderpoweredVotes = 0

		store.Set(iterator.Key(), k.cdc.MustMarshalBinaryBare(resetCard))
	}
}

func (k Keeper) GetCardsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.cardsStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
