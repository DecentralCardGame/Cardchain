package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgWithCreator interface {
	GetSigners() []sdk.AccAddress
}

const MAX_ALIAS_LEN = 25

// User Combines types.User and it's account address for better usability
type User struct {
	types.User
	Addr sdk.AccAddress
}

// CreateUser Creates a new user
func (k Keeper) CreateUser(ctx sdk.Context, addr sdk.AccAddress, alias string) error {
	// check if user already exists
	_, err := k.GetUser(ctx, addr)
	if err != nil {
		err = k.InitUser(ctx, addr, alias)
		if err != nil {
			return err
		}
	} else {
		return types.ErrUserAlreadyExists
	}
	return nil
}

// InitUser Initializes a new user
func (k Keeper) InitUser(ctx sdk.Context, address sdk.AccAddress, alias string) error {
	if alias == "" {
		alias = "newbie"
	}
	newUser := types.NewUser()
	err := checkAliasLimit(alias)
	if err != nil {
		return err
	}
	newUser.Alias = alias
	err = k.MintCoinsToAddr(ctx, address, sdk.Coins{sdk.NewInt64Coin("ucredits", 10000000000)})
	if err != nil {
		return err
	}

	matchesEnabled, err := k.FeatureFlagModuleInstance.Get(ctx, string(types.FeatureFlagName_Matches))
	if err != nil {
		return err
	}
	if !matchesEnabled {
		newUser.VotableCards = k.GetAllVotableCards(ctx)
	}

	userObj := User{newUser, address}
	k.ClaimAirDrop(ctx, &userObj, types.AirDrop_user)
	k.SetUserFromUser(ctx, userObj)
	return nil
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

func (k Keeper) GetMsgCreator(ctx sdk.Context, msg MsgWithCreator) (user User, err error) {
	user.Addr = msg.GetSigners()[0]

	user.User, err = k.GetUser(ctx, user.Addr)
	return
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
		allAddresses = append(allAddresses, iterator.Key())
	}
	return
}

func checkAliasLimit(alias string) error {
	if len(alias) > MAX_ALIAS_LEN {
		return sdkerrors.Wrapf(types.InvalidAlias, "alias is too long (%d) maximum is %d", len(alias), MAX_ALIAS_LEN)
	}

	return nil
}
