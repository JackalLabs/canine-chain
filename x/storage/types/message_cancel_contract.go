package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelContract = "cancel_contract"

var _ sdk.Msg = &MsgCancelContract{}

func NewMsgCancelContract(creator string, cid string) *MsgCancelContract {
	return &MsgCancelContract{
		Creator: creator,
		Cid:     cid,
	}
}

func (msg *MsgCancelContract) Route() string {
	return RouterKey
}

func (msg *MsgCancelContract) Type() string {
	return TypeMsgCancelContract
}

func (msg *MsgCancelContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
