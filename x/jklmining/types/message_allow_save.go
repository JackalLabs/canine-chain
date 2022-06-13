package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAllowSave = "allow_save"

var _ sdk.Msg = &MsgAllowSave{}

func NewMsgAllowSave(creator string, passkey string, size string) *MsgAllowSave {
	return &MsgAllowSave{
		Creator: creator,
		Passkey: passkey,
		Size_:   size,
	}
}

func (msg *MsgAllowSave) Route() string {
	return RouterKey
}

func (msg *MsgAllowSave) Type() string {
	return TypeMsgAllowSave
}

func (msg *MsgAllowSave) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAllowSave) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAllowSave) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
