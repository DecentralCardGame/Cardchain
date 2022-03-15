package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math/big"
)

// MulCoin multiplies a Coin with an int
func MulCoin(coin sdk.Coin, amt int64) sdk.Coin {
	return sdk.Coin{coin.Denom, coin.Amount.Mul(sdk.NewInt(amt))}
}

// MulCoinFloat multiplies a Coin with a float
func MulCoinFloat(coin sdk.Coin, amt float64) sdk.Coin {
	amount := big.NewFloat(amt)
	oldAmount := new(big.Float).SetInt(coin.Amount.BigInt())
	oldAmount.Mul(amount, oldAmount)
	var newAmount big.Int
	oldAmount.Int(&newAmount)
	return sdk.Coin{coin.Denom, sdk.NewIntFromBigInt(&newAmount)}
}

// QuoCoin devides a Coin with by int
func QuoCoin(coin sdk.Coin, amt int64) sdk.Coin {
	return sdk.Coin{coin.Denom, coin.Amount.Quo(sdk.NewInt(amt))}
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

// MintCoinsToString adds coins to an Account
func (k Keeper) MintCoinsToString(ctx sdk.Context, user string, amounts sdk.Coins) error {
	addr, err := sdk.AccAddressFromBech32(user)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}
	err = k.MintCoinsToAddr(ctx, addr, amounts)
	if err != nil {
		return err
	}
	return nil
}

// BurnCoinsFromString adds coins to an Account
func (k Keeper) BurnCoinsFromString(ctx sdk.Context, user string, amounts sdk.Coins) error {
	addr, err := sdk.AccAddressFromBech32(user)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidAccAddress, "Unable to convert to AccAddress")
	}
	err = k.BurnCoinsFromAddr(ctx, addr, amounts)
	if err != nil {
		return err
	}
	return nil
}

// TransferFromCoin Transfers coins from a certain pool to a coin
func (k Keeper) TransferFromCoin(ctx sdk.Context, addr string, pool *sdk.Coin, coin sdk.Coin) error {
	err := k.MintCoinsToString(ctx, addr, sdk.Coins{coin})
	if err != nil {
		return err
	}
	*pool = pool.Sub(coin)
	return nil
}
