package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetProviderTotalspace = "set_provider_totalspace"

var _ sdk.Msg = &MsgSetProviderTotalspace{}

func NewMsgSetProviderTotalspace(creator string, space string) *MsgSetProviderTotalspace {
	return &MsgSetProviderTotalspace{
		Creator: creator,
		Space:   space,
	}
}

func (msg *MsgSetProviderTotalspace) Route() string {
	return RouterKey
}

func (msg *MsgSetProviderTotalspace) Type() string {
	return TypeMsgSetProviderTotalspace
}

func (msg *MsgSetProviderTotalspace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProviderTotalspace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProviderTotalspace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
