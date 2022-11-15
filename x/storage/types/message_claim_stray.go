package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
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
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != addressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Cid)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid (%s)", err)
	}
	if prefix != cidPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jklc`", prefix))
	}
	return nil
}
