package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSignContract = "sign_contract"

var _ sdk.Msg = &MsgSignContract{}

func NewMsgSignContract(creator string, cid string, payOnce bool) *MsgSignContract {
	return &MsgSignContract{
		Creator: creator,
		Cid:     cid,
		PayOnce: payOnce,
	}
}

func (msg MsgSignContract) Route() string {
	return RouterKey
}

func (msg MsgSignContract) Type() string {
	return TypeMsgSignContract
}

func (msg MsgSignContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg MsgSignContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSignContract) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Cid)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid (%s)", err)
	}
	if prefix != CidPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jklc`", prefix))
	}
	return nil
}
