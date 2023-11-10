package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProvisionFileTree = "provision_file_tree"

var _ sdk.Msg = &MsgProvisionFileTree{}

func NewMsgProvisionFileTree(creator string, editors string, viewers string, trackingNumber string) *MsgProvisionFileTree {
	return &MsgProvisionFileTree{
		Creator:        creator,
		Editors:        editors,
		Viewers:        viewers,
		TrackingNumber: trackingNumber,
	}
}

func (msg *MsgProvisionFileTree) Route() string {
	return RouterKey
}

func (msg *MsgProvisionFileTree) Type() string {
	return TypeMsgProvisionFileTree
}

func (msg *MsgProvisionFileTree) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProvisionFileTree) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProvisionFileTree) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Editors == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid editors: %s", msg.Editors)
	}
	if msg.Viewers == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid viewers: %s", msg.Viewers)
	}
	if msg.TrackingNumber == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid tracking number: %s", msg.TrackingNumber)
	}

	return nil
}
