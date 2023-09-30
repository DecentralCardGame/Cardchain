package cardchain

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgAddArtwork:
			res, err := msgServer.AddArtwork(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBuyCardScheme:
			res, err := msgServer.BuyCardScheme(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgVoteCard:
			res, err := msgServer.VoteCard(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSaveCardContent:
			res, err := msgServer.SaveCardContent(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgTransferCard:
			res, err := msgServer.TransferCard(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDonateToCard:
			res, err := msgServer.DonateToCard(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateuser:
			res, err := msgServer.Createuser(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgChangeArtist:
			res, err := msgServer.ChangeArtist(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRegisterForCouncil:
			res, err := msgServer.RegisterForCouncil(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgReportMatch:
			res, err := msgServer.ReportMatch(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		// case *types.MsgApointMatchReporter: // Will be uncommented later when I know how to check for module account
		// 	res, err := msgServer.ApointMatchReporter(sdk.WrapSDKContext(ctx), msg)
		// 	return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateSet:
			res, err := msgServer.CreateSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddCardToSet:
			res, err := msgServer.AddCardToSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgFinalizeSet:
			res, err := msgServer.FinalizeSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBuySet:
			res, err := msgServer.BuySet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveCardFromSet:
			res, err := msgServer.RemoveCardFromSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveContributorFromSet:
			res, err := msgServer.RemoveContributorFromSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddContributorToSet:
			res, err := msgServer.AddContributorToSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSubmitSetProposal:
			res, err := msgServer.SubmitSetProposal(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateSellOffer:
			res, err := msgServer.CreateSellOffer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBuyCard:
			res, err := msgServer.BuyCard(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveSellOffer:
			res, err := msgServer.RemoveSellOffer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddArtworkToSet:
			res, err := msgServer.AddArtworkToSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddStoryToSet:
			res, err := msgServer.AddStoryToSet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetCardRarity:
			res, err := msgServer.SetCardRarity(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateCouncil:
			res, err := msgServer.CreateCouncil(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCommitCouncilResponse:
			res, err := msgServer.CommitCouncilResponse(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRevealCouncilResponse:
			res, err := msgServer.RevealCouncilResponse(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRestartCouncil:
			res, err := msgServer.RestartCouncil(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRewokeCouncilRegistration:
			res, err := msgServer.RewokeCouncilRegistration(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgConfirmMatch:
			res, err := msgServer.ConfirmMatch(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetProfileCard:
			res, err := msgServer.SetProfileCard(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgOpenBoosterPack:
			res, err := msgServer.OpenBoosterPack(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgTransferBoosterPack:
			res, err := msgServer.TransferBoosterPack(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetSetStoryWriter:
			res, err := msgServer.SetSetStoryWriter(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetSetArtist:
			res, err := msgServer.SetSetArtist(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetUserWebsite:
			res, err := msgServer.SetUserWebsite(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetUserBiography:
			res, err := msgServer.SetUserBiography(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(errors.ErrUnknownRequest, errMsg)
		}
	}
}

// apply nerf levels and remove inappropriate cards

func UpdateNerfLevels(ctx sdk.Context, keeper keeper.Keeper) sdk.Result {

	buffbois, nerfbois, _, banbois := keeper.GetOPandUPCards(ctx)
	keeper.NerfBuffCards(ctx, buffbois, true)
	keeper.NerfBuffCards(ctx, nerfbois, false)
	keeper.UpdateBanStatus(ctx, banbois)

	keeper.ResetAllVotes(ctx)
	keeper.RemoveExpiredVoteRights(ctx)

	return sdk.Result{}
}
