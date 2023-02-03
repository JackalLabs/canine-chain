package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddClaimer = "add_claimer"
const TypeMsgRemoveClaimer = "remove_claimer"

var _ sdk.Msg = &MsgAddClaimer{}
var _ sdk.Msg = &MsgRemoveClaimer{}

func NewMsgAddClaimer(creator string, claimer string) *MsgAddClaimer {
	return &MsgAddClaimer{
		Creator:      creator,
		ClaimAddress: claimer,
	}
}

func (msg *MsgAddClaimer) Route() string {
	return RouterKey
}

func (msg *MsgAddClaimer) Type() string {
	return TypeMsgAddClaimer
}

func (msg *MsgAddClaimer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddClaimer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddClaimer) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.ClaimAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}

func NewMsgRemoveClaimer(creator string, claimer string) *MsgRemoveClaimer {
	return &MsgRemoveClaimer{
		Creator:      creator,
		ClaimAddress: claimer,
	}
}

func (msg *MsgRemoveClaimer) Route() string {
	return RouterKey
}

func (msg *MsgRemoveClaimer) Type() string {
	return TypeMsgRemoveClaimer
}

func (msg *MsgRemoveClaimer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveClaimer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveClaimer) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.ClaimAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
