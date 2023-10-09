package keeper

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QCard(goCtx context.Context, req *types.QueryQCardRequest) (*types.OutpCard, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Start of query code
	cardId, err := strconv.ParseUint(req.CardId, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(errors.ErrUnknownRequest, "could not parse cardId")
	}

	card := k.Cards.Get(ctx, cardId)
	if card == nil {
		return nil, sdkerrors.Wrap(errors.ErrUnknownRequest, "cardId does not represent a card")
	}

	image := k.Images.Get(ctx, card.ImageId)
	sum := md5.Sum(image.Image)
	hash := hex.EncodeToString(sum[:])

	outpCard := types.OutpCard{
		Owner:              card.Owner,
		Artist:             card.Artist,
		Content:            string(card.Content),
		Image:              string(image.Image),
		FullArt:            card.FullArt,
		Notes:              card.Notes,
		Status:             card.Status,
		VotePool:           card.VotePool,
		Voters:             card.Voters,
		FairEnoughVotes:    card.FairEnoughVotes,
		OverpoweredVotes:   card.OverpoweredVotes,
		UnderpoweredVotes:  card.UnderpoweredVotes,
		InappropriateVotes: card.InappropriateVotes,
		Nerflevel:          card.Nerflevel,
		BalanceAnchor:      card.BalanceAnchor,
		Hash:               hash,
  Rarity:             card.Rarity,
	}

	return &outpCard, nil
}
