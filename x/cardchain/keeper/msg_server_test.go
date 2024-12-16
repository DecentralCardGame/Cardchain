package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DecentralCardGame/cardchain/testutil/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CardchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
