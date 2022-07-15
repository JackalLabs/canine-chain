package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMinerTotalspace = "set_miner_totalspace"

var _ sdk.Msg = &MsgSetMinerTotalspace{}

func NewMsgSetMinerTotalspace(creator string, space string) *MsgSetMinerTotalspace {
	return &MsgSetMinerTotalspace{
		Creator: creator,
		Space:   space,
	}
}

func (msg *MsgSetMinerTotalspace) Route() string {
	return RouterKey
}

func (msg *MsgSetMinerTotalspace) Type() string {
	return TypeMsgSetMinerTotalspace
}

func (msg *MsgSetMinerTotalspace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMinerTotalspace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMinerTotalspace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
