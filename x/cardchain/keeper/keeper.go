package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec // The wire codec for binary encoding/decoding.
		UsersStoreKey   sdk.StoreKey // UsersStoreKey
		CardsStoreKey     sdk.StoreKey // Cardstorekey
		InternalStoreKey sdk.StoreKey
		paramstore paramtypes.Subspace

		BankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	usersStoreKey,
	cardsStoreKey sdk.StoreKey,
	internalStoreKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:              cdc,
		UsersStoreKey:    usersStoreKey,
		CardsStoreKey:    cardsStoreKey,
		InternalStoreKey: internalStoreKey,
		paramstore:       ps,
		BankKeeper:       bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Adds coins to an Account
func (k Keeper) MintCoinsToAddr(ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	coinMint := types.CoinsIssuerName

	// mint coins to minter module account
	err := k.BankKeeper.MintCoins(ctx, coinMint, amounts)
	if err != nil {
		return err
	}

	// send coins to the address
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, coinMint, addr, amounts)
	if err != nil {
		return err
	}

	return nil
}

// Removes Coins from an Account
func (k Keeper) BurnCoinsFromAddr(ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	coinMint := types.CoinsIssuerName

	// send coins to the module
	err := k.BankKeeper.SendCoinsFromAccountToModule(ctx, addr, coinMint, amounts)
	if err != nil {
		return err
	}

	// burn coins from minter module account
	err = k.BankKeeper.BurnCoins(ctx, coinMint, amounts)
	if err != nil {
		return err
	}

	return nil
}

// GetLastCardScheme - get the current id of the last bought card scheme
func (k Keeper) GetLastCardSchemeId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("lastCardScheme"))
	return binary.BigEndian.Uint64(bz)
}

// SetLastCardScheme - sets the current id of the last bought card scheme
func (k Keeper) SetLastCardSchemeId(ctx sdk.Context, lastId uint64) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("lastCardScheme"), sdk.Uint64ToBigEndian(lastId))
}

// returns the current price of the card scheme auction
func (k Keeper) GetCardAuctionPrice(ctx sdk.Context) sdk.Coin {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
	var price sdk.Coin
	k.cdc.MustUnmarshal(bz, &price)
	return price
}

// adds or subtracts credits from the public pool
func (k Keeper) AddPublicPoolCredits(ctx sdk.Context, delta sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshal(bz, &amount)
	newAmount := amount.Add(delta)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshal(&newAmount))
}

// sets the current price of the card scheme auction
func (k Keeper) SetCardAuctionPrice(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("currentCardSchemeAuctionPrice"), k.cdc.MustMarshal(&price))
}

func (k Keeper) SetCard(ctx sdk.Context, cardId uint64, newCard types.Card) {
	store := ctx.KVStore(k.CardsStoreKey)
	store.Set(sdk.Uint64ToBigEndian(cardId), k.cdc.MustMarshal(&newCard))
}

func (k Keeper) AddOwnedCardScheme(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)

	var gottenUser types.User
	k.cdc.MustUnmarshal(bz, &gottenUser)

	gottenUser.OwnedCardSchemes = append(gottenUser.OwnedCardSchemes, cardId)

	store.Set(address, k.cdc.MustMarshal(&gottenUser))
}

func (k Keeper) Createuser(ctx sdk.Context, newUser sdk.AccAddress, alias string) types.User {
	// check if user already exists
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(newUser)

	if bz == nil {
		k.InitUser(ctx, newUser, alias)
	}
	return k.GetUser(ctx, newUser)
}

func (k Keeper) InitUser(ctx sdk.Context, address sdk.AccAddress, alias string) {
	store := ctx.KVStore(k.UsersStoreKey)

	if alias == "" {
		alias = "newbie"
	}
	newUser := types.NewUser()
	newUser.Alias = alias
	k.MintCoinsToAddr(ctx, address, sdk.Coins{sdk.NewInt64Coin("credits", 10000)})
	const votingRightsExpirationTime = 86000
	newUser.VoteRights = k.GetVoteRightToAllCards(ctx, ctx.BlockHeight()+votingRightsExpirationTime) // TODO this might be a good thing to remove later, so that sybil voting is not possible

	store.Set(address, k.cdc.MustMarshal(&newUser))
}

func (k Keeper) GetUser(ctx sdk.Context, address sdk.AccAddress) types.User {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)

	var gottenUser types.User
	k.cdc.MustUnmarshal(bz, &gottenUser)
	return gottenUser
}

func (k Keeper) GetVoteRightToAllCards(ctx sdk.Context, expireBlock int64) []*types.VoteRight {
	cardStore := ctx.KVStore(k.CardsStoreKey)
	cardIterator := sdk.KVStorePrefixIterator(cardStore, nil)

	votingRights := []*types.VoteRight{}

	for ; cardIterator.Valid(); cardIterator.Next() {
		// here only give right if card is not a scheme or banished
		var gottenCard types.Card
		k.cdc.MustUnmarshal(cardIterator.Value(), &gottenCard)

		if gottenCard.Status == "permanent" || gottenCard.Status == "trial" || gottenCard.Status == "prototype" {
			right := types.NewVoteRight(binary.BigEndian.Uint64(cardIterator.Key()), expireBlock)
			votingRights = append(votingRights, &right)
		}
	}
	cardIterator.Close()

	return votingRights
}
