package keeper

import (
    "context"

    "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (k Keeper) QAccountFromZealy(goCtx context.Context, req *types.QueryQAccountFromZealyRequest) (*types.QueryQAccountFromZealyResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    ctx := sdk.UnwrapSDKContext(goCtx)

    zealy, err := k.GetZealy(ctx, req.ZealyId)
    if err != nil {
        return nil, err
    }

    return &types.QueryQAccountFromZealyResponse{Address: zealy.Address}, nil
}
