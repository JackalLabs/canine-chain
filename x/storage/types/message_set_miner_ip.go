package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetProviderIP = "set_provider_ip"

var _ sdk.Msg = &MsgSetProviderIP{}

func NewMsgSetProviderIP(creator string, ip string) *MsgSetProviderIP {
	return &MsgSetProviderIP{
		Creator: creator,
		Ip:      ip,
	}
}

func (msg *MsgSetProviderIP) Route() string {
	return RouterKey
}

func (msg *MsgSetProviderIP) Type() string {
	return TypeMsgSetProviderIP
}

func (msg *MsgSetProviderIP) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProviderIP) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProviderIP) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
