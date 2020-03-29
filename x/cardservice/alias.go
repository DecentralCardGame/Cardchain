package cardservice

import (
	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	CardsStoreKey     = types.StoreKey
	UsersStoreKey     = types.StoreKey
	InternalStoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper        			= keeper.NewKeeper
	NewQuerier       			= keeper.NewQuerier
	NewMsgBuyCardScheme 	= types.NewMsgBuyCardScheme
	NewMsgSaveCardContent = types.NewMsgSaveCardContent
	NewMsgVoteCard				= types.NewMsgVoteCard
	NewMsgTransferCard 		= types.NewMsgTransferCard
	NewMsgDonateToCard		= types.NewMsgDonateToCard
	NewMsgCreateUser			=	types.NewMsgCreateUser

	ModuleCdc        			= types.ModuleCdc
	RegisterCodec    			= types.RegisterCodec
)

type (
	Keeper          		= keeper.Keeper
	MsgBuyCardScheme 		= types.MsgBuyCardScheme
	MsgSaveCardContent  = types.MsgSaveCardContent
	MsgVoteCard					= types.MsgVoteCard
	MsgTransferCard 		= types.MsgTransferCard
	MsgDonateToCard			= types.MsgDonateToCard
	MsgCreateUser				=	types.MsgCreateUser
	User								= types.User
	Card								= types.Card
	VoteRight						= types.VoteRight

)
