package types

import (
	fmt "fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEncounterCreate = "encounter_create"

var _ sdk.Msg = &MsgEncounterCreate{}

func NewMsgEncounterCreate(creator string, name string, drawlist []uint64, parameters []*Parameter, image []byte) *MsgEncounterCreate {
	return &MsgEncounterCreate{
		Creator:    creator,
		Name:       name,
		Drawlist:   drawlist,
		Parameters: parameters,
		Image:      image,
	}
}

func (msg *MsgEncounterCreate) Route() string {
	return RouterKey
}

func (msg *MsgEncounterCreate) Type() string {
	return TypeMsgEncounterCreate
}

func (msg *MsgEncounterCreate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEncounterCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEncounterCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Drawlist) > 40 || len(msg.Drawlist) < 1 {
		return sdkerrors.Wrapf(ErrInvalidData, "invalid drawlist length, max 40 is '%d'", len(msg.Drawlist))
	}

	if msg.Name == "" {
		return sdkerrors.Wrap(ErrInvalidData, "encounter needs a name")
	}

	if len(msg.Image) > ArtworkMaxSize {
		return sdkerrors.Wrap(ErrImageSizeExceeded, fmt.Sprint(len(msg.Image)))
	}

	return nil
}
