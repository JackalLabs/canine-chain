package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddViewers = "add_viewers"

var _ sdk.Msg = &MsgAddViewers{}

func NewMsgAddViewers(creator string, viewerIds string, viewerKeys string, address string, owner string, viewersToNotify string, notiForViewers string) *MsgAddViewers {
	return &MsgAddViewers{
		Creator:        creator,
		ViewerIds:      viewerIds,
		ViewerKeys:     viewerKeys,
		Address:        address,
		Fileowner:      owner,
		NotifyViewers:  viewersToNotify,
		NotiForViewers: notiForViewers,
	}
}

func (msg *MsgAddViewers) Route() string {
	return RouterKey
}

func (msg *MsgAddViewers) Type() string {
	return TypeMsgAddViewers
}

func (msg *MsgAddViewers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddViewers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddViewers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
