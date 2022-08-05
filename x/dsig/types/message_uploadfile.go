package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUploadfile = "uploadfile"

var _ sdk.Msg = &MsgUploadfile{}

func NewMsgUploadfile(creator string, fid string) *MsgUploadfile {
	return &MsgUploadfile{
		Creator: creator,
		Fid:     fid,
	}
}

func (msg *MsgUploadfile) Route() string {
	return RouterKey
}

func (msg *MsgUploadfile) Type() string {
	return TypeMsgUploadfile
}

func (msg *MsgUploadfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUploadfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUploadfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
