package types

import (
	fmt "fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetProviderTotalspace = "set_provider_totalspace"

var _ sdk.Msg = &MsgSetProviderTotalspace{}

func NewMsgSetProviderTotalspace(creator string, space string) *MsgSetProviderTotalspace {
	return &MsgSetProviderTotalspace{
		Creator: creator,
		Space:   space,
	}
}

func (msg *MsgSetProviderTotalspace) Route() string {
	return RouterKey
}

func (msg *MsgSetProviderTotalspace) Type() string {
	return TypeMsgSetProviderTotalspace
}

func (msg *MsgSetProviderTotalspace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProviderTotalspace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProviderTotalspace) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	if _, err := strconv.Atoi(msg.Space); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse space (%s)", err)
	}
	return nil
}
