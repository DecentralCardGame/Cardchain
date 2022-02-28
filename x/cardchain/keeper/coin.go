package keeper

import (
  "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
  sdk "github.com/cosmos/cosmos-sdk/types"
)

// MulCoin multiplies a Coin with an int
func MulCoin(coin sdk.Coin, amt int64) sdk.Coin {
  return sdk.Coin{coin.Denom, coin.Amount.Mul(sdk.NewInt(amt))}
}

// MintCoinsToAddr adds coins to an Account
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

//  BurnCoinsFromAddr removes Coins from an Account
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
