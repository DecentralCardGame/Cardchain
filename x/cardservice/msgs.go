package cardservice

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

//////////////
// Set Type //
//////////////

// MsgSetName defines a SetName message
type MsgSetType struct {
	NameID string
	Value  string
	Owner  sdk.AccAddress
}

// NewSetNameMsg is a constructor function for MsgSetType
func NewMsgSetType(name string, value string, owner sdk.AccAddress) MsgSetType {
	return MsgSetType{
		NameID: name,
		Value:  value,
		Owner:  owner,
	}
}

// Name Implements Msg.
func (msg MsgSetType) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgSetType) Type() string { return "set_type" }

// ValdateBasic Implements Msg.
func (msg MsgSetType) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.NameID) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetType) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSetType) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

///////////////
// Vote Card //
///////////////

// MsgVoteCard defines a VoteCard message
type MsgVoteCard struct {
	CardID string
	Type   string
	Voter  sdk.AccAddress
}

// NewMsgVoteCard is a constructor function for MsgVoteCard
func NewMsgVoteCard(cardId string, type string, voter sdk.AccAddress) MsgSetType {
	return MsgSetType{
		CardID: cardId,
		Type:   type,
		Voter:  voter,
	}
}

// Name Implements Msg.
func (msg MsgVoteCard) Route() string { return "cardservice" }

// Type Implements Msg.
func (msg MsgVoteCard) Type() string { return "vote_card" }

// ValdateBasic Implements Msg.
func (msg MsgVoteCard) ValidateBasic() sdk.Error {
	if msg.MsgVoteCard.Empty() {
		return sdk.ErrInvalidAddress(msg.MsgVoteCard.String())
	}
	if len(msg.CardID) == 0 || len(msg.Type) == 0 {
		return sdk.ErrUnknownRequest("CardId and/or Vote Typeui cannot be empty")
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
