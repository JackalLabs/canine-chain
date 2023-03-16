package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgResetEditors = "reset_editors"

var _ sdk.Msg = &MsgResetEditors{}

func NewMsgResetEditors(creator string, address string, fileowner string) *MsgResetEditors {
	return &MsgResetEditors{
		Creator:   creator,
		Address:   address,
		Fileowner: fileowner,
	}
}

func (msg *MsgResetEditors) Route() string {
	return RouterKey
}

func (msg *MsgResetEditors) Type() string {
	return TypeMsgResetEditors
}

func (msg *MsgResetEditors) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgResetEditors) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResetEditors) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid address: %s", msg.Address)
	}
	if msg.Fileowner == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid file owner: %s", msg.Fileowner)
	}

	return nil
}
