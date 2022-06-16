package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimSave = "claim_save"

var _ sdk.Msg = &MsgClaimSave{}

func NewMsgClaimSave(creator string, saveindex string, key string) *MsgClaimSave {
	return &MsgClaimSave{
		Creator:   creator,
		Saveindex: saveindex,
		Key:       key,
	}
}

func (msg *MsgClaimSave) Route() string {
	return RouterKey
}

func (msg *MsgClaimSave) Type() string {
	return TypeMsgClaimSave
}

func (msg *MsgClaimSave) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimSave) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimSave) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
