package keeper_test

import (
	"context"
	"testing"

	keepertest "cardchain/testutil/keeper"
	"cardchain/x/cardchain/keeper"
	"cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CardchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
