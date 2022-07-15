package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSignContract = "sign_contract"

var _ sdk.Msg = &MsgSignContract{}

func NewMsgSignContract(creator string, cid string) *MsgSignContract {
	return &MsgSignContract{
		Creator: creator,
		Cid:     cid,
	}
}

func (msg *MsgSignContract) Route() string {
	return RouterKey
}

func (msg *MsgSignContract) Type() string {
	return TypeMsgSignContract
}

func (msg *MsgSignContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSignContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
