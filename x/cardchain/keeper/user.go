package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// User Combines types.User and it's account address for better usability
type User struct {
	*types.User
	Addr sdk.AccAddress
}

// InitUser Initializes a new user
func (k Keeper) InitUser(ctx sdk.Context, address sdk.AccAddress, alias string) error {
	if alias == "" {
		alias = "newbie"
	}
	newUser := types.NewUser()

	newUser.Alias = alias
	err := k.MintCoinsToAddr(ctx, address, sdk.Coins{sdk.NewInt64Coin("ucredits", 10000000000)})
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

	userObj := User{&newUser, address}
	k.ClaimAirDrop(ctx, &userObj, types.AirDrop_user)
	k.SetUserFromUser(ctx, userObj)
	return nil
}

// GetUserFromString Gets a user from store, but instead of taking an accountaddress it takes a string and returns a User struct (defined above)
func (k Keeper) GetUserFromString(ctx sdk.Context, addr string) (user User, err error) {
	user.Addr, err = sdk.AccAddressFromBech32(addr)
	if err != nil {
		return user, sdkerrors.Wrap(errors.ErrInvalidAddress, "Unable to convert to AccAddress")
	}
	user.User = k.Users.Get(ctx, user.Addr)
	if err != nil {
		return
	}
	return
}

// SetUserFromUser Sets a user in store, but takes a User struct as defined above
func (k Keeper) SetUserFromUser(ctx sdk.Context, user User) {
	k.Users.Set(ctx, user.Addr, user.User)
}

// GetAllUsers Gets all users from store
func (k Keeper) GetAllUsers(ctx sdk.Context) (allUsers []*types.User, allAddresses []sdk.AccAddress) {
	iterator := k.Users.GetIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenUser types.User
		k.cdc.MustUnmarshal(iterator.Value(), &gottenUser)

		allUsers = append(allUsers, &gottenUser)
		allAddresses = append(allAddresses, iterator.Key())
	}
	return
}
