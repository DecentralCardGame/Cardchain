package cardservice

import (
	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/types"
)

const (
	ModuleName 	= types.ModuleName
	RouterKey  	= types.RouterKey
	StoreKey		= types.StoreKey
	Cards   		= types.StoreKey
	Users				= types.StoreKey
	Internal		= types.StoreKey
)

var (
	NewKeeper        			= keeper.NewKeeper
	NewQuerier       			= keeper.NewQuerier
	NewMsgBuyCardScheme		= types.NewMsgBuyCardScheme
	NewMsgSaveCardContent	= types.NewMsgSaveCardContent
	NewMsgVoteCard 				= types.NewMsgVoteCard
	NewMsgTransferCard 		= types.NewMsgTransferCard
	NewMsgDonateToCard 		= types.NewMsgDonateToCard
	NewMsgCreateUser 			= types.NewMsgCreateUser
	NewUser         			= types.NewUser
	NewCard								= types.NewCard
	ModuleCdc        			= types.ModuleCdc
	RegisterCodec    			= types.RegisterCodec
)

type (
	Keeper      		    = keeper.Keeper
	MsgBuyCardScheme    = types.MsgBuyCardScheme
	MsgSaveCardContent	= types.MsgSaveCardContent
	MsgVoteCard   			= types.MsgVoteCard
	MsgTransferCard 		= types.MsgTransferCard
	MsgDonateToCard   	= types.MsgDonateToCard
	MsgCreateUser				= types.MsgCreateUser
	Card								= types.Card
	User								= types.User
	VoteRights					= types.VoteRight
)
