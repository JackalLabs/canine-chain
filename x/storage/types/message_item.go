package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgItem = "item"

var _ sdk.Msg = &MsgItem{}

func NewMsgItem(creator string, hashlist string) *MsgItem {
	return &MsgItem{
		Creator:  creator,
		Hashlist: hashlist,
	}
}

func (msg *MsgItem) Route() string {
	return RouterKey
}

func (msg *MsgItem) Type() string {
	return TypeMsgItem
}

func (msg *MsgItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
