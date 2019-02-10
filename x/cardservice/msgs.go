package cardservice

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


type User struct {
	voteRights []VoteRight
}


type Card struct {
	Owner sdk.AccAddress
	Content []byte
	Status string
	VotePool sdk.Coin
	FairEnoughVotes uint64
	OverpoweredVotes uint64
	UnderpoweredVotes uint64
	InappropriateVotes uint64
	Nerflevel int64
}

func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner: owner,
		Content: []byte{},
		Status: "scheme",
		VotePool: sdk.NewInt64Coin("credits", 0),
		FairEnoughVotes: 0,
		OverpoweredVotes: 0,
		UnderpoweredVotes: 0,
		InappropriateVotes: 0,
		Nerflevel: 0,
	}
}
/*
// MsgSetName defines a SetName message
type MsgSetName struct {
	NameID string
	Value  string
	Owner  sdk.AccAddress
}

// NewSetNameMsg is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		NameID: name,
		Value:  value,
		Owner:  owner,
	}
}

// Name Implements Msg.
func (msg MsgSetName) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgSetName) Type() string { return "set_card" }

// ValdateBasic Implements Msg.
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.NameID) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetName) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
*/
/////////////////////
// Buy Card Scheme //
/////////////////////

// MsgBuyName defines the BuyName message
type MsgBuyCardScheme struct {
	//NameID string
	Bid    sdk.Coin
	Buyer  sdk.AccAddress
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyCardScheme(bid sdk.Coin, buyer sdk.AccAddress) MsgBuyCardScheme {
	return MsgBuyCardScheme{
		//NameID: name,
		Bid:    bid,
		Buyer:  buyer,
	}
}

// Name Implements Msg.
func (msg MsgBuyCardScheme) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgBuyCardScheme) Type() string { return "buy_card_scheme" }

// ValidateBasic Implements Msg.
func (msg MsgBuyCardScheme) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	/*if len(msg.NameID) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}*/
	if !msg.Bid.IsPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBuyCardScheme) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgBuyCardScheme) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

///////////////////////
// Save Card Content //
///////////////////////

// MsgBuyName defines the BuyName message
type MsgSaveCardContent struct {
	CardId	uint64
	Content []byte
	Owner		sdk.AccAddress
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgSaveCardContent(cardId uint64, content []byte, owner sdk.AccAddress) MsgSaveCardContent {
	return MsgSaveCardContent{
		CardId: 	cardId,
		Content:	content,
		Owner: 		owner,
	}
}

// Name Implements Msg.
func (msg MsgSaveCardContent) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgSaveCardContent) Type() string { return "save_card_content" }

// ValidateBasic Implements Msg.
func (msg MsgSaveCardContent) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Content) == 0 {
		return sdk.ErrUnknownRequest("Content cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSaveCardContent) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSaveCardContent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

///////////////
// Vote Card //
///////////////

// MsgVoteCard defines a VoteCard message
type MsgVoteCard struct {
	CardID 		uint64
	VoteType	string
	Voter  		sdk.AccAddress
}

// NewMsgVoteCard is a constructor function for MsgVoteCard
func NewMsgVoteCard(cardId uint64, voteType string, voter sdk.AccAddress) MsgVoteCard {
	return MsgVoteCard{
		CardID:		cardId,
		VoteType:	voteType,
		Voter:		voter,
	}
}

// Name Implements Msg.
func (msg MsgVoteCard) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgVoteCard) Type() string { return "vote_card" }

// ValdateBasic Implements Msg.
func (msg MsgVoteCard) ValidateBasic() sdk.Error {
	if msg.Voter.Empty() {
		return sdk.ErrInvalidAddress(msg.Voter.String())
	}
	// the check of CardID < 0 might be pointless.. should be validated in the rest api or nscli
	if msg.CardID < 0 || len(msg.VoteType) == 0 {
		return sdk.ErrUnknownRequest("CardId and/or Vote Type cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgVoteCard) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgVoteCard) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Voter}
}

///////////////////
// Transfer Card //
///////////////////

// MsgTransferCard defines a TransferCard message
type MsgTransferCard struct {
	CardID 		uint64
	Source 		sdk.AccAddress
	Target		sdk.AccAddress
}

// NewMsgTransferCard is a constructor function for MsgTransferCard
func NewMsgTransferCard(cardId uint64, source sdk.AccAddress, target sdk.AccAddress) MsgTransferCard {
	return MsgTransferCard{
		CardID:		cardId,
		Source:		source,
		Target:		target,
	}
}

// Name Implements Msg.
func (msg MsgTransferCard) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgTransferCard) Type() string { return "vote_card" }

// ValdateBasic Implements Msg.
func (msg MsgTransferCard) ValidateBasic() sdk.Error {
	if msg.Source.Empty() {
		return sdk.ErrInvalidAddress(msg.Source.String())
	}
	if msg.Target.Empty() {
		return sdk.ErrInvalidAddress(msg.Target.String())
	}
	// the check of CardID < 0 might be pointless.. should be validated in the rest api or nscli
	if msg.CardID < 0 {
		return sdk.ErrUnknownRequest("CardId and/or Vote Type cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransferCard) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}


////////////////////
// Donate to Card //
////////////////////

// MsgDonateToCard defines a TransferCard message
type MsgDonateToCard struct {
	CardID 		uint64
	Donator 	sdk.AccAddress
	Amount		sdk.Coin
}

// NewMsgMsgDonateToCard is a constructor function for MsgDonateToCard
func NewMsgDonateToCard(cardId uint64, donator sdk.AccAddress, amount sdk.Coin) MsgDonateToCard {
	return MsgDonateToCard{
		CardID:		cardId,
		Donator:	donator,
		Amount:		amount,
	}
}

// Name Implements Msg.
func (msg MsgDonateToCard) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgDonateToCard) Type() string { return "donate_to_card" }

// ValdateBasic Implements Msg.
func (msg MsgDonateToCard) ValidateBasic() sdk.Error {
	if msg.Donator.Empty() {
		return sdk.ErrInvalidAddress(msg.Donator.String())
	}
	// the check of CardID < 0 might be pointless.. should be validated in the rest api or nscli
	if msg.CardID < 0 || msg.Amount.IsZero() {
		return sdk.ErrUnknownRequest("CardId cannot be empty and Amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDonateToCard) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
