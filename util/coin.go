package util

import (
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MulCoin multiplies a Coin with an int
func MulCoin(coin sdk.Coin, amt int64) sdk.Coin {
	return sdk.Coin{
		Denom:  coin.Denom,
		Amount: coin.Amount.Mul(math.NewInt(amt)),
	}
}

// MulCoinFloat multiplies a Coin with a float
func MulCoinFloat(coin sdk.Coin, amt float64) sdk.Coin {
	amount := big.NewFloat(amt)
	oldAmount := new(big.Float).SetInt(coin.Amount.BigInt())
	oldAmount.Mul(amount, oldAmount)
	var newAmount big.Int
	oldAmount.Int(&newAmount)
	return sdk.Coin{
		Denom:  coin.Denom,
		Amount: math.NewIntFromBigInt(&newAmount),
	}
}

// QuoCoin devides a Coin with by int
func QuoCoin(coin sdk.Coin, amt int64) sdk.Coin {
	return sdk.Coin{
		Denom:  coin.Denom,
		Amount: coin.Amount.Quo(math.NewInt(amt)),
	}
}
