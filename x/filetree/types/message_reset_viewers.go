package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgResetViewers = "reset_viewers"

var _ sdk.Msg = &MsgResetViewers{}

func NewMsgResetViewers(creator string, address string, fileowner string) *MsgResetViewers {
	return &MsgResetViewers{
		Creator:   creator,
		Address:   address,
		Fileowner: fileowner,
	}
}

func (msg *MsgResetViewers) Route() string {
	return RouterKey
}

func (msg *MsgResetViewers) Type() string {
	return TypeMsgResetViewers
}

func (msg *MsgResetViewers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgResetViewers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResetViewers) ValidateBasic() error {
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
