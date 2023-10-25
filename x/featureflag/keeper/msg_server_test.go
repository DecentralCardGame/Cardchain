package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DecentralCardGame/Cardchain/testutil/keeper"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/keeper"
	"github.com/DecentralCardGame/Cardchain/x/featureflag/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FeatureflagKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
