package keeper

import (
	"context"
	"fmt"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountFromZealy(goCtx context.Context, req *types.QueryAccountFromZealyRequest) (*types.QueryAccountFromZealyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	zealy := k.Zealy.Get(ctx, req.ZealyId)
	if zealy == nil {
		return nil, fmt.Errorf("zealyId `%s` not in store", req.ZealyId)
	}

	return &types.QueryAccountFromZealyResponse{Address: zealy.Address}, nil
}
