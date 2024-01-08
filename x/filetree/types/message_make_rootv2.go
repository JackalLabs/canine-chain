package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

const TypeMsgMakeRootV2 = "make_root_v2"

var (
	_ sdk.Msg            = &MsgMakeRootV2{}
	_ legacytx.LegacyMsg = &MsgMakeRootV2{}
)

func NewMsgMakeRootV2(creator string, editors string, viewers string, trackingNumber string) *MsgMakeRootV2 {
	return &MsgMakeRootV2{
		Creator:        creator,
		Editors:        editors,
		Viewers:        viewers,
		TrackingNumber: trackingNumber,
	}
}

func (msg *MsgMakeRootV2) Route() string {
	return RouterKey
}

func (msg *MsgMakeRootV2) Type() string {
	return TypeMsgMakeRootV2
}

func (msg *MsgMakeRootV2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMakeRootV2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMakeRootV2) ValidateBasic() error {
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
