package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	bpf  = "bpf"  // 1 (base denom unit)
	mbpf = "mbpf" // 10^-3 (milli)
	ubpf = "ubpf" // 10^-6 (micro)
)

func RegisterNativeCoinUnits() {
	_ = sdk.RegisterDenom(bpf, sdk.OneDec())
	_ = sdk.RegisterDenom(mbpf, sdk.NewDecWithPrec(1, 3))
	_ = sdk.RegisterDenom(ubpf, sdk.NewDecWithPrec(1, 6))
}
