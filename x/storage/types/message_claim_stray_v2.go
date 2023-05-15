package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimStrayV2 = "claim_stray_v2"

var _ sdk.Msg = &MsgClaimStrayV2{}

func NewMsgClaimStrayV2(creator string, cid string, forAddress string) *MsgClaimStrayV2 {
	return &MsgClaimStrayV2{
		Creator:    creator,
		Cid:        cid,
		ForAddress: forAddress,
	}
}

func (msg *MsgClaimStrayV2) Route() string {
	return RouterKey
}

func (msg *MsgClaimStrayV2) Type() string {
	return TypeMsgClaimStrayV2
}

func (msg *MsgClaimStrayV2) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimStrayV2) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimStrayV2) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.ForAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid for address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid for prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Cid)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid (%s)", err)
	}
	if prefix != CidPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jklc`", prefix))
	}
	return nil
}
