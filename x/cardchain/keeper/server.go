package keeper

import (
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetServerId Gets the number of a certain server by it's reporters address
func (k Keeper) GetServerId(ctx sdk.Context, addr string) uint64 {
	for idx, server := range k.servers.GetAll(ctx) {
		if server.Reporter == addr {
			return uint64(idx)
		}
	}
	id := k.servers.GetNumber(ctx)
	server := types.Server{Reporter: addr}
	k.servers.Set(ctx, id, &server)
	return id
}

// ReportServerMatch Adds a number of reports to the server record
func (k Keeper) ReportServerMatch(ctx sdk.Context, addr string, num uint64, isValid bool) {
	id := k.GetServerId(ctx, addr)
	server := k.servers.Get(ctx, id)
	switch isValid {
	case true:
		server.ValidReports += num
	case false:
		server.InvalidReports += num
	}
	k.servers.Set(ctx, id, server)
}
