package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitAll = "init_all"

var _ sdk.Msg = &MsgInitAll{}

func NewMsgInitAll(creator string, pubkey string) *MsgInitAll {
	return &MsgInitAll{
		Creator: creator,
		Pubkey:  pubkey,
	}
}

func (msg *MsgInitAll) Route() string {
	return RouterKey
}

func (msg *MsgInitAll) Type() string {
	return TypeMsgInitAll
}

func (msg *MsgInitAll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitAll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitAll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
