package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChoosePlan = "choose_plan"

var _ sdk.Msg = &MsgChoosePlan{}

func NewMsgChoosePlan(creator string, tbCount string, denom string) *MsgChoosePlan {
	return &MsgChoosePlan{
		Creator:      creator,
		TbCount:      tbCount,
		PaymentDenom: denom,
	}
}

func (msg *MsgChoosePlan) Route() string {
	return RouterKey
}

func (msg *MsgChoosePlan) Type() string {
	return TypeMsgChoosePlan
}

func (msg *MsgChoosePlan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChoosePlan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChoosePlan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
