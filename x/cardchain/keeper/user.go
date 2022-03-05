package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// User Combines types.User and it's account address for better usability
type User struct {
	types.User
	Addr sdk.AccAddress
}

// CreateUser Creates a new user
func (k Keeper) CreateUser(ctx sdk.Context, newUser sdk.AccAddress, alias string) types.User {
	// check if user already exists
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(newUser)

	if bz == nil {
		k.InitUser(ctx, newUser, alias)
	}

	user, _ := k.GetUser(ctx, newUser)
	return user
}

// InitUser Initializes a new user
func (k Keeper) InitUser(ctx sdk.Context, address sdk.AccAddress, alias string) {
	store := ctx.KVStore(k.UsersStoreKey)

	if alias == "" {
		alias = "newbie"
	}
	newUser := types.NewUser()
	newUser.Alias = alias
	k.MintCoinsToAddr(ctx, address, sdk.Coins{sdk.NewInt64Coin("ucredits", 10000000000)})
	newUser.VoteRights = k.GetVoteRightToAllCards(ctx, ctx.BlockHeight()+k.GetParams(ctx).VotingRightsExpirationTime) // TODO this might be a good thing to remove later, so that sybil voting is not possible

	store.Set(address, k.cdc.MustMarshal(&newUser))
}

// GetUser Gets a user from store
func (k Keeper) GetUser(ctx sdk.Context, address sdk.AccAddress) (types.User, error) {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)
	var gottenUser types.User

	if bz == nil {
		return gottenUser, sdkerrors.Wrap(types.ErrUserDoesNotExist, "Address not in store")
	}

	k.cdc.MustUnmarshal(bz, &gottenUser)
	return gottenUser, nil
}

// GetUserFromString Gets a user from store, but instead of taking an accountaddress it takes a string and returns a User struct (defined above)
func (k Keeper) GetUserFromString(ctx sdk.Context, addr string) (user User, err error) {
	user.Addr, err = sdk.AccAddressFromBech32(addr)
	if err != nil {
		return user, sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}
	user.User, err = k.GetUser(ctx, user.Addr)
	if err != nil {
		return
	}
	return
}

// SetUser Sets a user in store
func (k Keeper) SetUser(ctx sdk.Context, address sdk.AccAddress, userData types.User) {
	store := ctx.KVStore(k.UsersStoreKey)
	store.Set(address, k.cdc.MustMarshal(&userData))
}

// SetUserFromUser Sets a user in store, but takes a User struct as defined above
func (k Keeper) SetUserFromUser(ctx sdk.Context, user User) {
	k.SetUser(ctx, user.Addr, user.User)
}

// GetUsersIterator Returns an iterator for all users
func (k Keeper) GetUsersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.UsersStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// GetAllUsers Gets all users from store
func (k Keeper) GetAllUsers(ctx sdk.Context) (allUsers []*types.User, allAddresses []sdk.AccAddress) {
	iterator := k.GetUsersIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenUser types.User
		k.cdc.MustUnmarshal(iterator.Value(), &gottenUser)

		allUsers = append(allUsers, &gottenUser)
		allAddresses = append(allAddresses, sdk.AccAddress(iterator.Key()))
	}
	return
}
