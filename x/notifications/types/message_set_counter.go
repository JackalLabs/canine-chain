package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCounter = "set_counter"

var _ sdk.Msg = &MsgSetCounter{}

func NewMsgSetCounter(creator string) *MsgSetCounter {
	return &MsgSetCounter{
		Creator: creator,
	}
}

func (msg *MsgSetCounter) Route() string {
	return RouterKey
}

func (msg *MsgSetCounter) Type() string {
	return TypeMsgSetCounter
}

func (msg *MsgSetCounter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCounter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCounter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
