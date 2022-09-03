package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRecord = "add_record"

var _ sdk.Msg = &MsgAddRecord{}

func NewMsgAddRecord(creator string, name string, record string, value string, data string) *MsgAddRecord {
	return &MsgAddRecord{
		Creator: creator,
		Name:    name,
		Value:   value,
		Data:    data,
		Record:  record,
	}
}

func (msg *MsgAddRecord) Route() string {
	return RouterKey
}

func (msg *MsgAddRecord) Type() string {
	return TypeMsgAddRecord
}

func (msg *MsgAddRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
