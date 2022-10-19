package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimStray = "claim_stray"

var _ sdk.Msg = &MsgClaimStray{}

func NewMsgClaimStray(creator string, cid string) *MsgClaimStray {
	return &MsgClaimStray{
		Creator: creator,
		Cid:     cid,
	}
}

func (msg *MsgClaimStray) Route() string {
	return RouterKey
}

func (msg *MsgClaimStray) Type() string {
	return TypeMsgClaimStray
}

func (msg *MsgClaimStray) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimStray) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimStray) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
