package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddViewers = "add_viewers"

var _ sdk.Msg = &MsgAddViewers{}

func NewMsgAddViewers(creator string, viewerIds string, viewerKeys string, address string, owner string) *MsgAddViewers {
	return &MsgAddViewers{
		Creator:    creator,
		ViewerIds:  viewerIds,
		ViewerKeys: viewerKeys,
		Address:    address,
		Fileowner:  owner,
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
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("`%s` is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Fileowner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Fileowner address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Fileowner prefix (%s)", fmt.Errorf("`%s` is not a valid prefix here. Expected `jkl`", prefix))
	}
	return nil
}
