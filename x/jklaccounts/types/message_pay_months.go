package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPayMonths = "pay_months"

var _ sdk.Msg = &MsgPayMonths{}

func NewMsgPayMonths(creator string, address string, months string, paymentDenom string) *MsgPayMonths {
	return &MsgPayMonths{
		Creator:      creator,
		Address:      address,
		Months:       months,
		PaymentDenom: paymentDenom,
	}
}

func (msg *MsgPayMonths) Route() string {
	return RouterKey
}

func (msg *MsgPayMonths) Type() string {
	return TypeMsgPayMonths
}

func (msg *MsgPayMonths) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPayMonths) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPayMonths) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
