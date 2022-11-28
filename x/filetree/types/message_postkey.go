package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostkey = "postkey"

var _ sdk.Msg = &MsgPostkey{}

func NewMsgPostkey(creator string, key string) *MsgPostkey {
	return &MsgPostkey{
		Creator: creator,
		Key:     key,
	}
}

func (msg *MsgPostkey) Route() string {
	return RouterKey
}

func (msg *MsgPostkey) Type() string {
	return TypeMsgPostkey
}

func (msg *MsgPostkey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostkey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostkey) ValidateBasic() error {
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
