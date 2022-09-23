package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitAccount = "init_account"

var _ sdk.Msg = &MsgInitAccount{}

func NewMsgInitAccount(creator string, account string, rootHashpath string, editors string, key string, trackingNumber uint64) *MsgInitAccount {
	return &MsgInitAccount{
		Creator:        creator,
		Account:        account,
		RootHashpath:   rootHashpath,
		Editors:        editors,
		Key:            key,
		TrackingNumber: trackingNumber,
	}
}

func (msg *MsgInitAccount) Route() string {
	return RouterKey
}

func (msg *MsgInitAccount) Type() string {
	return TypeMsgInitAccount
}

func (msg *MsgInitAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
