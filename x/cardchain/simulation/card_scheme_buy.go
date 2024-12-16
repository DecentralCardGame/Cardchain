package simulation

import (
	"math/rand"

	"github.com/DecentralCardGame/cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCardSchemeBuy(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCardSchemeBuy{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CardSchemeBuy simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CardSchemeBuy simulation not implemented"), nil, nil
	}
}
