package types

import (
	fmt "fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEncounterCreate = "encounter_create"

var _ sdk.Msg = &MsgEncounterCreate{}

func NewMsgEncounterCreate(creator string, drawlist []uint64, parameters map[string]string, image []byte) *MsgEncounterCreate {
	return &MsgEncounterCreate{
		Creator:    creator,
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

	if len(msg.Drawlist) != 40 {
		return sdkerrors.Wrapf(ErrInvalidData, "drawlist too long, must be 40 is '%d'", len(msg.Drawlist))
	}

	if len(msg.Image) > ArtworkMaxSize {
		return sdkerrors.Wrap(ErrImageSizeExceeded, fmt.Sprint(len(msg.Image)))
	}

	return nil
}
