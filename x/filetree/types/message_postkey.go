package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostKey = "post_key"

var _ sdk.Msg = &MsgPostKey{}

func NewMsgPostKey(creator string, key string) *MsgPostKey {
	return &MsgPostKey{
		Creator: creator,
		Key:     key,
	}
}

func (msg *MsgPostKey) Route() string {
	return RouterKey
}

func (msg *MsgPostKey) Type() string {
	return TypeMsgPostKey
}

func (msg *MsgPostKey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostKey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Key == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid key: %s", msg.Key)
	}

	return nil
}
