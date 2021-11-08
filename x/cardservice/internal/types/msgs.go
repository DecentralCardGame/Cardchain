package types

import (
	"strconv"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/////////////////////
// Buy Card Scheme //
/////////////////////

// MsgBuyCard defines the BuyCard message
type MsgBuyCardScheme struct {
	Bid    sdk.Coin				`json:"bid"`
	Buyer  sdk.AccAddress `json:"buyer"`
}

// NewMsgBuyCard is the constructor function for MsgBuyCard
func NewMsgBuyCardScheme(bid sdk.Coin, buyer sdk.AccAddress) MsgBuyCardScheme {
	return MsgBuyCardScheme{
		Bid:    bid,
		Buyer:  buyer,
	}
}

// Route should return the name of the module
func (msg MsgBuyCardScheme) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyCardScheme) Type() string { return "buy_card_scheme" }

// ValidateBasic Implements Msg.
func (msg MsgBuyCardScheme) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}
	if !msg.Bid.IsPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyCardScheme) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyCardScheme) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

///////////////////////
// Save Card Content //
///////////////////////

// MsgBuyName defines the BuyName message
type MsgSaveCardContent struct {
	CardId	uint64					`json:"cardid"`
	Content []byte					`json:"content"`
	Image 	[]byte					`json:"image"`
	FullArt string					`json:"fullart"`
	Notes 	string					`json:"notes"`
	Owner		sdk.AccAddress	`json:"owner"`
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgSaveCardContent(cardId uint64, content []byte, image []byte, fullart string, notes string, owner sdk.AccAddress) MsgSaveCardContent {
	return MsgSaveCardContent{
		CardId: 	cardId,
		Content:	content,
		Image: 		image,
		FullArt:  fullart,
		Notes:		notes,
		Owner: 		owner,
	}
}

// Name Implements Msg.
func (msg MsgSaveCardContent) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSaveCardContent) Type() string { return "save_card_content" }

// ValidateBasic Implements Msg.
func (msg MsgSaveCardContent) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Content) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Content cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveCardContent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSaveCardContent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

///////////////
// Vote Card //
///////////////

// MsgVoteCard defines a VoteCard message
type MsgVoteCard struct {
	CardId 		uint64					`json:"cardid"`
	VoteType	string					`json:"votetype"`
	Voter  		sdk.AccAddress	`json:"voter"`
}

// NewMsgVoteCard is a constructor function for MsgVoteCard
func NewMsgVoteCard(cardId uint64, voteType string, voter sdk.AccAddress) MsgVoteCard {
	return MsgVoteCard{
		CardId:		cardId,
		VoteType:	voteType,
		Voter:		voter,
	}
}

// Route should return the name of the module
func (msg MsgVoteCard) Route() string { return RouterKey }

// Type should return the action
func (msg MsgVoteCard) Type() string { return "vote_card" }

// ValdateBasic Implements Msg.
func (msg MsgVoteCard) ValidateBasic() error {
	if msg.Voter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Voter.String())
	}
	// the check of CardID < 0 might be pointless.. should be validated in the rest api or cscli
	if msg.CardId < 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "CardId is: "+strconv.FormatUint(msg.CardId, 10)+" - should be non-negative")
	}
	if len(msg.VoteType) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Vote Type is: "+msg.VoteType+" - cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgVoteCard) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgVoteCard) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Voter}
}

///////////////////
// Transfer Card //
///////////////////

// MsgTransferCard defines a TransferCard message
type MsgTransferCard struct {
	CardId 		uint64					`json:"cardid"`
	Sender 		sdk.AccAddress	`json:"sender"`
	Receiver	sdk.AccAddress	`json:"receiver"`
}

// NewMsgTransferCard is a constructor function for MsgTransferCard
func NewMsgTransferCard(cardId uint64, sender sdk.AccAddress, receiver sdk.AccAddress) MsgTransferCard {
	return MsgTransferCard{
		CardId:		cardId,
		Sender:		sender,
		Receiver:	receiver,
	}
}

func (msg MsgTransferCard) Route() string { return RouterKey }

func (msg MsgTransferCard) Type() string { return "transfer_card" }

// ValdateBasic Implements Msg.
func (msg MsgTransferCard) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender.String())
	}
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Receiver.String())
	}
	// the check of CardID < 0 might be pointless.. should be validated in the rest api or cscli ?
	if msg.CardId < 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "CardId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransferCard) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgTransferCard) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

////////////////////
// Donate to Card //
////////////////////

// MsgDonateToCard defines a TransferCard message
type MsgDonateToCard struct {
	CardId 		uint64					`json:"cardid"`
	Donator 	sdk.AccAddress	`json:"donator"`
	Amount		sdk.Coin				`json:"amount"`
}

// NewMsgMsgDonateToCard is a constructor function for MsgDonateToCard
func NewMsgDonateToCard(cardId uint64, donator sdk.AccAddress, amount sdk.Coin) MsgDonateToCard {
	return MsgDonateToCard{
		CardId:		cardId,
		Donator:	donator,
		Amount:		amount,
	}
}

func (msg MsgDonateToCard) Route() string { return RouterKey }

func (msg MsgDonateToCard) Type() string { return "donate_to_card" }

// ValdateBasic Implements Msg.
func (msg MsgDonateToCard) ValidateBasic() error {
	if msg.Donator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Donator.String())
	}
	// the check of CardID < 0 might be pointless.. should be validated in the rest api or cscli ?
	if msg.CardId < 0 || msg.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "CardId cannot be empty or donation cannot be 0.")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDonateToCard) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgDonateToCard) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Donator}
}

/////////////////
// Create User //
/////////////////

// MsgCreateUser defines a CreateUser message
type MsgCreateUser struct {
	NewUser 	sdk.AccAddress	`json:"newuser"`
	Creator 	sdk.AccAddress	`json:"creator"`
	Alias			string					`json:"alias"`
}

// NewMsgCreateUser is a constructor function for MsgCreateUser
func NewMsgCreateUser(creator sdk.AccAddress, newUser sdk.AccAddress, alias string) MsgCreateUser {
	return MsgCreateUser{
		NewUser:	newUser,
		Creator: creator,
		Alias: alias,
	}
}

func (msg MsgCreateUser) Route() string { return RouterKey }

func (msg MsgCreateUser) Type() string { return "create_user" }

// ValdateBasic Implements Msg.
func (msg MsgCreateUser) ValidateBasic() error {
	if msg.NewUser.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.NewUser.String())
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCreateUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgCreateUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}
