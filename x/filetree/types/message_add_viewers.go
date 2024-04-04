package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddViewers = "add_viewers"

var _ sdk.Msg = &MsgAddViewers{}

func NewMsgAddViewers(creator string, viewerIDs string, viewerKeys string, address string, owner string) *MsgAddViewers {
	return &MsgAddViewers{
		Creator:    creator,
		ViewerIds:  viewerIDs,
		ViewerKeys: viewerKeys,
		Address:    address,
		FileOwner:  owner,
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

	// Check empty values
	if msg.ViewerIds == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid viewer id: %s", msg.ViewerIds)
	}
	if msg.ViewerKeys == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid viewer keys: %s", msg.ViewerKeys)
	}
	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid address: %s", msg.Address)
	}
	if msg.FileOwner == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid file owner: %s", msg.FileOwner)
	}

	return nil
}
