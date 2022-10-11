package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMakeFolder = "make_folder"

var _ sdk.Msg = &MsgMakeFolder{}

func NewMsgMakeFolder(creator string, account string, rootHashPath string, contents string, editors string, viewers string, trackingNumber uint64) *MsgMakeFolder {
	return &MsgMakeFolder{
		Creator:        creator,
		Account:        account,
		RootHashPath:   rootHashPath,
		Contents:       contents,
		Editors:        editors,
		Viewers:        viewers,
		TrackingNumber: trackingNumber,
	}
}

func (msg *MsgMakeFolder) Route() string {
	return RouterKey
}

func (msg *MsgMakeFolder) Type() string {
	return TypeMsgMakeFolder
}

func (msg *MsgMakeFolder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMakeFolder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMakeFolder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
